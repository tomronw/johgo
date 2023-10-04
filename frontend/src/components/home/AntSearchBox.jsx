import React from 'react';
import { Input } from 'antd';
import { toast } from 'sonner';
import '../../css/SearchBox.css'
import { useContext } from 'react';
import SearchContext from "../../context/SearchContext";
import { AutoComplete } from 'antd';
import TrendingUpIcon from '@mui/icons-material/TrendingUp';
import { useTheme } from '../../context/ThemeContext';
import { ConfigProvider, theme } from 'antd';
import SuggestionsContext from '../../context/SuggestionsContext';


function AntSearchBox({ checked }) {

    const { fetchProducts } = useContext(SearchContext)
    const { suggestions } = useContext(SuggestionsContext)
    const { darkMode } = useTheme();
    const isMobile = window.innerWidth <= 768;
    var currentTheme

    const renderTitle = (title) => (
        <span>
            {title}
        </span>
    );
    const renderItem = (title) => ({
        value: title,
        label: (
            <div
                style={{
                    display: 'flex',
                    justifyContent: 'space-between',
                }}
            >
                {title}
                <span>
                    <TrendingUpIcon />
                </span>
            </div>
        ),
    });

    const options = [
        {
            label: renderTitle('All trending searches'),
            options: suggestions.map(suggestion => renderItem(suggestion)),
        },
    ];


    const handleSearch = (value) => {
        let searchValueToUse = value;
        if (value && value.target) {
            searchValueToUse = value.target.value;
        }
        if (searchValueToUse === '') {
            toast.error('Please enter a search query.');
        } else if (searchValueToUse.toLowerCase() === 'pokemon' || searchValueToUse.length < 2) {
            toast.error('Please try to be more specific! I.E. "booster box" or "elite trainer box" ');
        } else {
            fetchProducts(searchValueToUse, checked);
        }
    };
    if (darkMode) {
        currentTheme = theme.darkAlgorithm
    } else {
        currentTheme = theme.lightAlgorithm
    }

    return (
        <div>
            <ConfigProvider
                theme={{
                    algorithm: currentTheme,
                }}
            >
                {isMobile ? (
                    <Input.Search
                        size='medium'
                        placeholder="Search JohGo..."
                        allowClear
                        bordered
                        onSearch={handleSearch}
                        onPressEnter={handleSearch}
                        style={{
                            width: '100%',
                            fontWeight: 600,
                            opacity: 0.6,
                        }}
                    />
                ) : (
                    <AutoComplete
                        popupClassName="certain-category-search-dropdown"
                        options={options}
                        size="large"
                        style={{
                            width: '80vh',
                            fontWeight: 600,
                            opacity: 0.6,
                        }}
                    >
                        <Input.Search
                            size='large'
                            placeholder="Search JohGo..."
                            allowClear
                            bordered
                            onSearch={handleSearch}
                            onPressEnter={handleSearch}
                            style={{
                                width: '100%',
                                fontWeight: 600,
                                opacity: 0.6,
                            }}
                        />
                    </AutoComplete>
                )}
            </ConfigProvider>
        </div>
    )

}

export default AntSearchBox;
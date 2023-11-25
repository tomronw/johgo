import React, { useState, useRef, useContext } from 'react';
import '../../../../css/SearchBox.css';
import { toast } from 'sonner';
import SearchContext from "../../../../context/SearchContext";
import { Input } from 'antd';
import { useTheme } from '../../../../context/ThemeContext.js';
import { ConfigProvider, theme } from 'antd';

function SearchBar() {
    const [isExpanded, setExpanded] = useState(false);
    const searchBoxRef = useRef(null);
    const { fetchProducts } = useContext(SearchContext)
    const [searchValue, setSearchValue] = useState('');
    const { darkMode } = useTheme();

    var currentTheme

    const handleInputChange = (event) => {
        setSearchValue(event.target.value);
    };

    if (darkMode) {
        currentTheme = theme.darkAlgorithm
    } else {
        currentTheme = theme.lightAlgorithm
    }

    const handleKeyDown = (event) => {
        if (event.keyCode === 13) {
            if (searchValue === '') {
                toast.error('Please enter a search query.');
            } else if (searchValue.toLowerCase() === 'pokemon') {
                toast.error('Please try to be more specific! I.E. "booster box" or "elite trainer box" ');

            } else {
                fetchProducts(searchValue, true)
            }
        }
    };

    const handleMouseEnter = () => {
        setExpanded(true);
    };

    const handleMouseLeave = () => {
        setExpanded(false);
    };

    return (
        <ConfigProvider
            theme={{
                algorithm: currentTheme,
            }}
        >
            <div className="searchBarContainer">
                <div
                    ref={searchBoxRef}
                    className={`searchBoxIcon ${isExpanded ? "searchBoxExpanded" : ""}`}
                    onMouseEnter={handleMouseEnter}
                    onMouseLeave={handleMouseLeave}
                    style={{
                        width: isExpanded ? "50%" : "50px",
                        marginRight: isExpanded ? "2%" : "0px",
                        borderRadius: isExpanded ? "4px" : "50%",
                        marginTop: "-1%",
                    }}
                >
                    <Input.Search
                        size='large'
                        placeholder={isExpanded ? "Search JohGo..." : ""}
                        allowClear
                        bordered
                        onChange={handleInputChange}
                        onSearch={handleKeyDown}
                        onPressEnter={handleKeyDown}
                        style={{ width: '100%', fontWeight: 600, opacity: 0.6 }}
                    />
                </div>
            </div>
        </ConfigProvider>
    );
}

export default SearchBar;
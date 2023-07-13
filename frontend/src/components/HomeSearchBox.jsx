import React from 'react';
import { Autocomplete } from '@material-ui/lab';
import { TextField } from '@material-ui/core';
import "../css/SearchBox.css";
import { Toaster, toast } from 'sonner';
import { useContext } from 'react';
import SearchContext from "../context/SearchContext";
import { useFetchSearchSuggestions } from "./SearchSuggestions";
import TrendingBoxes from "./TrendingSearchesBox";
import CardsCheckbox from "./CardsCheckbox";



function HomeSearchBox() {
    const handleInputChange = (event, value) => {
        setSearchValue(value);
    };

    const SearchSuggestions = useFetchSearchSuggestions();

    const { fetchProducts } = useContext(SearchContext)

    const [searchValue, setSearchValue] = React.useState('');
    const [checked, setChecked] = React.useState(true);
    const handleChange = (event) => {
        setChecked(event.target.checked);
        //console.log(checked)
    };

    const isMobile = window.innerWidth <= 768;


    const handleKeyDown = (event) => {
        if (event.keyCode === 13) {

            if (searchValue === '') {
                toast.error('Please enter a search query.');
            } else if (searchValue.toLowerCase() === 'pokemon' || searchValue.length < 2) {
                toast.error('Please try to be more specific! I.E. "booster box" or "elite trainer box" ');
            } else {
                fetchProducts(searchValue, checked)
            }
        }
    };

    return (

        <div>
            <Toaster closeButton offset="50px" position="bottom-center" />

            <div style={{ display: 'flex', justifyContent: 'center', alignItems: 'center', color: 'white', opacity: 0.6, position: 'relative', paddingTop: isMobile ? '35%' : '18%' }}>
                <Autocomplete
                    freeSolo
                    options={SearchSuggestions}
                    renderInput={(params) => (
                        <TextField {...params}

                            label="Search JohGo"
                            variant="outlined"
                            InputProps={{
                                ...params.InputProps, style: { color: 'white', opacity: 1, fontWeight: 600 },
                                onKeyDown: handleKeyDown,
                            }} />
                    )}
                    className="searchBox"
                    onInputChange={handleInputChange}
                />
            </div>
            <div >
                {!isMobile ? <CardsCheckbox checked={checked} handleChange={handleChange} /> : null}
            </div>
            <div>
                {!isMobile ? <TrendingBoxes trendingTopics={SearchSuggestions} /> : null}
            </div>
        </div>
    );
}

export default HomeSearchBox;
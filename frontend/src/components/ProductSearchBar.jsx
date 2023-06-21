import React from 'react';
import { Autocomplete } from '@material-ui/lab';
import { TextField } from '@material-ui/core';
import "../css/SearchBox.css";
import { Toaster, toast } from 'sonner';
import { useContext } from 'react';
import SearchContext from "../context/SearchContext";
import { useFetchSearchSuggestions } from "./SearchSuggestions";


function ProductSearchBar() {
    const handleInputChange = (event, value) => {
        setSearchValue(value);
    };

    const SearchSuggestions = useFetchSearchSuggestions();

    const { fetchProducts } = useContext(SearchContext)

    const [searchValue, setSearchValue] = React.useState('');

    const handleKeyDown = (event) => {
        if (event.keyCode === 13) {
            if (searchValue === '') {
                toast.error('Please enter a search query.');
            } else if (searchValue.toLowerCase() === 'pokemon') {
                toast.error('Please try to be more specific! I.E. "booster box" or "elite trainer box" ');

            } else {
                // console.log('Will start search process with: ', searchValue)
                fetchProducts(searchValue)

            }
        }
    };


    return (<div>
        <Toaster closeButton offset="50px" position="bottom-center" />
        <div style={{ color: 'white', opacity: 0.6 }}>
            <Autocomplete
                freeSolo
                options={SearchSuggestions}
                renderInput={(params) => (
                    <TextField {...params}
                        label="Search JohGo again!"
                        variant="outlined"
                        size="small"
                        InputProps={{
                            ...params.InputProps, style: { color: 'white' },
                            onKeyDown: handleKeyDown
                        }}
                    />
                )}
                className="top-search-box"
                onInputChange={handleInputChange} />
        </div>
    </div>);
}

export default ProductSearchBar;
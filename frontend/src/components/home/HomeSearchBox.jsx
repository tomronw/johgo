import React, { useEffect, useContext } from 'react';
import { Grid } from '@mui/material';
import TrendingBoxes from './TrendingBoxes.jsx';
import ExcludeSingles from './ExcludeSingles.jsx';
import '../../css/SearchBox.css'
import AntSearchBox from './AntSearchBox.jsx';
import SuggestionsContext from '../../context/SuggestionsContext.js';

function HomeSearchBox() {
    const [checked, setChecked] = React.useState(true);
    const { fetchSuggestions } = useContext(SuggestionsContext)
    const isMobile = window.innerWidth <= 768;

    useEffect(() => {
        fetchSuggestions();
        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, []);

    const handleChange = (event) => {
        setChecked(event.target.checked);
    };

    return (
        <Grid container style={{ minHeight: '65vh' }} alignItems="center" justifyContent="center" className='search-elements'>
            <Grid item xs={12} md={6}>
                <div style={{ display: 'flex', flexDirection: 'column', alignItems: 'center' }}>
                    <ExcludeSingles checked={checked} handleChange={handleChange} isMobile={isMobile} />
                    <AntSearchBox
                        checked={checked}
                    />
                </div>
                <TrendingBoxes isMobile={isMobile} />
            </Grid>
        </Grid>
    );

}

export default HomeSearchBox;

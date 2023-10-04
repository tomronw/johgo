import React from 'react';
import { Grid } from '@mui/material';
import FilterDropdown from '../other/FilterDropdown';
import '../../../../css/MobileSearchBar.css'
import AntSearchBox from '../../../home/AntSearchBox';

function MobileProductToolbar({ resultsCount, filter, onFilterChange }) {

    return (
        <Grid container direction="column" alignItems="center" style={{ margin: 0, padding: 0 }}>
            <Grid item style={{ width: '100%' }}>
                <Grid container style={{ minHeight: '20vh', maxHeight: '15vh' }} >
                    <Grid item xs={12} md={6}>
                        <div style={{ display: 'flex', flexDirection: 'column', alignItems: 'center' }}>
                            <AntSearchBox checked />
                        </div>
                    </Grid>
                    <Grid item xs={12} md={6} style={{ display: 'flex', justifyContent: 'center', alignItems: 'center', marginTop: '-90px' }}>  {/* Adjusted marginTop */}
                        <div style={{ color: 'white', fontSize: '0.8em' }}>Displaying {resultsCount} results</div>
                    </Grid>
                </Grid>
            </Grid>
            <Grid item style={{ marginTop: '-50px' }}>
                <FilterDropdown
                    filter={filter}
                    onFilterChange={onFilterChange}
                    className="mobileDropdown" />
            </Grid>
        </Grid>




    );
}

export default MobileProductToolbar;

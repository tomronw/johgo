import React from 'react';
import ProductSearchBar from './ProductSearchBar';
import FilterDropdown from '../other/FilterDropdown';
import { Grid } from '@mui/material';
import '../../../../css/ProductGrid.css';
import ResultsCount from '../other/ResultsCount';

function ProductSearchToolbar({ resultsCount, filter, onFilterChange }) {

    return (
        <Grid container alignItems="center" spacing={3} className="search-toolbar" sx={{ paddingTop: '10%' }}>
            <Grid item xs={1} /> {/* Spacer */}

            <Grid item xs={2} align="left">
                <FilterDropdown
                    filter={filter}
                    onFilterChange={onFilterChange} />
            </Grid>

            <Grid item xs={6} align="center">
                <ProductSearchBar />
            </Grid>

            <Grid item xs={2} align="right">
                <ResultsCount resultsCount={resultsCount} />
            </Grid>

            <Grid item xs={1} /> {/* Spacer */}
        </Grid>
    );
}

export default ProductSearchToolbar;

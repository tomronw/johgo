import { useContext, useState, useMemo, useEffect } from "react";
import ProductItem from "./ProductItem";
import { Box, Grid } from '@mui/material';
import '../css/App.css';
import Loader from '../images/walking.gif'
import SearchContext from "../context/SearchContext";
import TopSearchBox from "./TopSearchBox";
import NoResults from '../images/noResults.png'
import InputLabel from '@mui/material/InputLabel';
import MenuItem from '@mui/material/MenuItem';
import Select from '@mui/material/Select';


function SearchResults() {
    const isMobile = window.innerWidth <= 768;
    let productsToLoad;

    if (isMobile) {
        productsToLoad = 10;
    } else {
        productsToLoad = 40;
    }

    const { products, loading } = useContext(SearchContext);
    const [sortOrder, setSortOrder] = useState('default');
    const [displayCount, setDisplayCount] = useState(productsToLoad);


    useEffect(() => {
        const handleScroll = () => {
            if (window.innerHeight + window.scrollY >= document.body.offsetHeight) {
                // User has scrolled to bottom, load more data
                setDisplayCount(displayCount + productsToLoad);
            }
        };
        // Add the event listener when component mounts and remove it when it unmounts
        window.addEventListener('scroll', handleScroll);
        return () => window.removeEventListener('scroll', handleScroll);
    }, [displayCount]);



    const sortedProducts = useMemo(() => {
        switch (sortOrder) {
            case 'low-to-high':
                return [...products].sort((a, b) => a.Price - b.Price);
            case 'high-to-low':
                return [...products].sort((a, b) => b.Price - a.Price);
            default:
                return products;
        }
    }, [products, sortOrder]);



    if (!loading) {

        return (
            <div style={{
                display: 'flex',
                justifyContent: 'center',
                alignItems: 'center'
            }}>
                <div className="product-grid-wrapper" style={{ display: 'flex', justifyContent: 'center' }}>
                    <div style={{ position: 'relative' }}>
                        <Box sx={{ position: 'fixed', zIndex: 1, width: '100%' }}>
                            <div className="search-box-position">
                                <TopSearchBox />
                            </div>
                            {products.length > 1 ? (
                                <div>
                                    <div className="sort-value">

                                        <InputLabel id="order-by-label" className="sort-input" style={{ color: 'white' }}>Sort by</InputLabel>
                                        <Select className="sort-input"
                                            fullWidth
                                            labelId="order-by"
                                            id="order-by-select"
                                            value={sortOrder}
                                            style={{ backgroundColor: '#222', color: 'white', borderRadius: '15px', opacity: 0.6, textAlign: 'left', maxWidth: '185px' }}
                                            label="Order"
                                            onChange={(e) => setSortOrder(e.target.value)}>
                                            <MenuItem value="default">Default</MenuItem>
                                            <MenuItem value="low-to-high">Price: Low to High</MenuItem>
                                            <MenuItem value="high-to-low">Price: High to Low</MenuItem>
                                        </Select>
                                    </div>
                                    <div>
                                        <p className="search-stats">Showing {products.length} results</p>
                                    </div>
                                </div>) : null}

                        </Box>
                    </div>

                    <Grid container spacing={4} sx={{ paddingY: 25, paddingX: 10, marginTop: -3 }} className="product-grid-item">
                        {products.length > 1 ? (
                            sortedProducts.slice(0, displayCount).map((product) => (
                                <Grid item xs={11} sm={6} md={2.3} key={product.Url}>
                                    <ProductItem key={product.id} product={product} />
                                </Grid>
                            ))
                        ) : (

                            <div class="no-results">
                                <p>No results found!</p>
                                <img src={NoResults} alt="No Results!" />
                            </div>

                        )}
                    </Grid>

                </div >
            </div>

        )

    } else {
        return <img src={Loader} alt="Loading..." className="loading-spinner" />
    }
}

export default SearchResults;
import { useContext, useState, useMemo, useEffect } from "react";
import ProductCard from "./ProductCard";
import { Grid } from '@mui/material';
import '../../css/App.css';
import Loader from '../../images/walking.gif'
import SearchContext from "../../context/SearchContext";
import NoResults from '../../images/noResults.png'
import ProductSearchToolbar from "./productSearchComponents/desktop/ProductSearchToolbar";
import MobileProductToolbar from "./productSearchComponents/mobile/MobileProductToolbar"

function ShowResults() {

    const isMobile = window.innerWidth <= 768;
    let productsToLoad;

    if (isMobile) {
        productsToLoad = 10;
    } else {
        productsToLoad = 40;
    }

    const { products, loading } = useContext(SearchContext);
    const [filter, setFilter] = useState('default');
    const [displayCount, setDisplayCount] = useState(productsToLoad);


    useEffect(() => {
        const handleScroll = () => {
            if (window.innerHeight + window.scrollY >= document.body.offsetHeight - 5) {
                setDisplayCount(prevCount => {
                    return prevCount + productsToLoad;
                });
            }
        };
        window.addEventListener('scroll', handleScroll);
        return () => window.removeEventListener('scroll', handleScroll);
    }, [productsToLoad, displayCount]);



    const sortedProducts = useMemo(() => {
        switch (filter) {
            case 'low-to-high':
                return [...products].sort((a, b) => a.price - b.price);
            case 'high-to-low':
                return [...products].sort((a, b) => b.price - a.price);
            case 'a-z':
                return [...products].sort((a, b) => a.title.localeCompare(b.title));
            case 'z-a':
                return [...products].sort((a, b) => b.title.localeCompare(a.title));
            default:
                return products;
        }
    }, [products, filter]);



    if (!loading) {
        if (isMobile) {
            return (
                <Grid container direction="column" >
                    <Grid item sx={{ marginTop: 0 }}>
                        <MobileProductToolbar
                            resultsCount={products.length}
                            filter={filter}
                            onFilterChange={setFilter} />
                    </Grid>
                    <Grid item>
                        <div className="product-grid-container">
                            <Grid
                                container
                                spacing={1}
                                justifyContent="center"
                                alignItems="center"
                                sx={{ paddingY: '12%', paddingX: '4%', marginTop: -3, right: 5 }}>
                                {products.length > 1 ? (
                                    sortedProducts.slice(0, displayCount).map((product) => (
                                        <Grid item xs={9} sm={6} md={4} lg={2} xl={1} key={product.url}>
                                            <ProductCard key={product.url} product={product} />
                                        </Grid>
                                    ))
                                ) : (
                                    <Grid item xs={12}>
                                        <div className="no-results-container">
                                            <h4>No results found!</h4>
                                            <img src={NoResults} alt="No Results!" />
                                        </div>
                                    </Grid>
                                )}
                            </Grid>
                        </div>
                    </Grid>
                </Grid>
            )
        } else {
            return (
                <Grid container direction="column" spacing={3}>
                    <Grid item>
                        <ProductSearchToolbar
                            resultsCount={products.length}
                            filter={filter}
                            onFilterChange={setFilter} />
                    </Grid>
                    <Grid item>
                        <div className="product-grid-container">
                            <Grid
                                container
                                spacing={4}
                                justifyContent="center"
                                alignItems="center"
                                sx={{ paddingY: '12%', paddingX: '4%', marginTop: -3, right: 5 }}
                            >
                                {products.length > 1 ? (
                                    sortedProducts.slice(0, displayCount).map((product) => (
                                        <Grid item xs={12} sm={6} md={4} lg={2} xl={1} key={product.url}>
                                            <ProductCard key={product.url} product={product} />
                                        </Grid>
                                    ))
                                ) : (
                                    <Grid item xs={12}>
                                        <div className="no-results-container">
                                            <h4>No results found!</h4>
                                            <img src={NoResults} alt="No Results!" />
                                        </div>
                                    </Grid>
                                )}
                            </Grid>
                        </div>
                    </Grid>
                </Grid>
            )
        }
    } else {
        return <img src={Loader} alt="Loading..." className="loading-spinner" />
    }
}

export default ShowResults;
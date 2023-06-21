import HomeSearchBox from "../components/HomeSearchBox.jsx";
import SearchResults from "../components/ShowResults.jsx";
import SearchContext from "../context/SearchContext";
import Loader from "../components/Loader.jsx";
import { useContext } from "react";
import '../css/App.css';


function Home() {

    const { products, loading } = useContext(SearchContext)


    if (!loading) {
        return (
            <div>
                <div className="App">
                    {products.length > 0 ? <SearchResults style={{ paddingTop: '60px' }} /> : <HomeSearchBox />}
                </div>
            </div>
        )
    } else {
        return (<div className="App" style={{ position: 'relative' }}>
            <Loader />
        </div>)

    }

}

export default Home;
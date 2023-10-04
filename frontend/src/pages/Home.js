import HomeSearchBox from "../components/home/HomeSearchBox.jsx";
import SearchContext from "../context/SearchContext.js";
import Loading from "../components/miscellaneous/Loading.jsx";
import { useContext } from "react";
import '../css/App.css';
import ShowResults from "../components/results/ShowResults.jsx";

function Home() {

    const { products, loading } = useContext(SearchContext)

    if (!loading) {
        return (
            <div>
                <div className="App">
                    {products.length > 0 ? <ShowResults /> : <HomeSearchBox />}
                </div>
            </div>
        )
    } else {
        return (<div className="App" style={{ position: 'relative' }}>
            <Loading />
        </div>)

    }

}

export default Home;
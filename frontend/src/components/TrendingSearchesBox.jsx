import { Chip } from '@material-ui/core';
import "../css/TrendingBox.css";
import TrendingUpIcon from '@mui/icons-material/TrendingUp';
import { useContext } from 'react';
import SearchContext from "../context/SearchContext";


function TrendingBoxes({ trendingTopics }) {

    const { fetchProducts } = useContext(SearchContext)


    const handleClick = (label) => {
        fetchProducts(label)
    };

    if (trendingTopics.length === 0) {
        return null;
    }
    return (
        <div className="trending">
            <Chip style={{ margin: '5px 10px' }} label={trendingTopics[0].toLowerCase()} onClick={() => handleClick(trendingTopics[0])} avatar={<TrendingUpIcon />} />
            <Chip style={{ margin: '5px 10px' }} label={trendingTopics[1].toLowerCase()} onClick={() => handleClick(trendingTopics[1])} avatar={<TrendingUpIcon />} />
            <Chip style={{ margin: '5px 10px' }} label={trendingTopics[2].toLowerCase()} onClick={() => handleClick(trendingTopics[2])} avatar={<TrendingUpIcon />} />
        </div>
    );
}

export default TrendingBoxes;
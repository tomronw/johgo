import { Chip } from '@material-ui/core';
import "../../css/Trending.css";
import TrendingUpIcon from '@mui/icons-material/TrendingUp';
import { useContext } from 'react';
import SearchContext from "../../context/SearchContext";
import SuggestionsContext from '../../context/SuggestionsContext';

function TrendingBoxes({ isMobile }) {

    const { fetchProducts } = useContext(SearchContext)
    const { suggestions } = useContext(SuggestionsContext)


    const handleClick = (label) => {
        fetchProducts(label, false)
    };

    return (
        <div>
            {!isMobile ? <div className="trending">
                <Chip style={{ margin: '15px 10px' }} label={suggestions[0].toLowerCase()} onClick={() => handleClick(suggestions[0])} avatar={<TrendingUpIcon />} />
                <Chip style={{ margin: '15px 10px' }} label={suggestions[1].toLowerCase()} onClick={() => handleClick(suggestions[1])} avatar={<TrendingUpIcon />} />
                <Chip style={{ margin: '15px 10px' }} label={suggestions[2].toLowerCase()} onClick={() => handleClick(suggestions[2])} avatar={<TrendingUpIcon />} /><br />
                <Chip style={{ margin: '5px 10px' }} label={suggestions[3].toLowerCase()} onClick={() => handleClick(suggestions[3])} avatar={<TrendingUpIcon />} />
                <Chip style={{ margin: '5px 10px' }} label={suggestions[4].toLowerCase()} onClick={() => handleClick(suggestions[4])} avatar={<TrendingUpIcon />} />
            </div> :
                <div className="trending">
                    <Chip style={{ margin: '10px 10px', fontSize: 'x-small' }} label={suggestions[0].toLowerCase()} onClick={() => handleClick(suggestions[2])} avatar={<TrendingUpIcon fontSize="small" />} /><br />
                    <Chip style={{ margin: '5px 10px', fontSize: 'x-small' }} label={suggestions[1].toLowerCase()} onClick={() => handleClick(suggestions[3])} avatar={<TrendingUpIcon fontSize="small" />} />
                </div>}
        </div>
    );
}

export default TrendingBoxes;
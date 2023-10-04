import Gif from '../../images/walking.gif'
import '../../css/App.css';
import '../../css/Index.css';

function Loading() {

    return (
        <img src={Gif} alt="Loading..." className="loading-spinner" />
    )
}

export default Loading;
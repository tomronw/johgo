import SMagikarp from '../images/magikarp.png';
import '../css/NotFound.css';

function NonExist() {
    return (<div>
        <img src={SMagikarp} alt='Magikarp' className='fish-flap'></img>
        <h1 className='not-found-text'>No page exists for the route: {window.location.pathname} </h1>
    </div>);
}

export default NonExist;
import SMagikarp from '../../images/magikarp.png';

function RouteNotFound() {
    return (<div>
        <img src={SMagikarp} alt='Magikarp' className='fish-flap'></img>
        <h1 style={{ color: 'white', fontSize: '17px' }}>No page exists for the route: {window.location.pathname} </h1>
    </div>);
}

export default RouteNotFound;
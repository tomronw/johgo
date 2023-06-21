import React from 'react';
import '../css/App.css';
import GitImage from '../images/git.png';
import Haunter from '../images/haunter.png';


function About() {
    return (

        <div>
            <div className="image-container">
                <img src={Haunter} alt="Haunter">
                </img>
            </div>
            <div className="black-box">
                <h2>About JohGo</h2>
                <p>No ads, no affiliate links. This is a completely free, open-source and community driven project for searching for in stock TCG products in the UK. </p>
            </div>
            <div className="contribute-box">
                <p>Developer/Engineer? Checkout the JohGo repo <a href="https://github.com/tomronw/johgo" rel="noreferrer" target='_blank' style={{ color: '#1d2bcc', textDecoration: 'none' }}> here!</a><img src={GitImage} alt="GitHub logo" className="github-logo" /></p>
            </div>
        </div >
    );
}

export default About;
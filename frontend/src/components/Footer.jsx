import React from 'react';
import '../css/Footer.css';
import Gastly from '../images/gastly.png'

function Footer() {
    const [showModal, setShowModal] = React.useState(false);

    const handleFairUseClick = () => {
        setShowModal(true);
    };

    const handleCloseModal = () => {
        setShowModal(false);
    };

    return (
        <div className="footer">
            <a href="https://github.com/tomronw" target="_blank" rel="noopener noreferrer">
                My Github
            </a>
            <span className="footer-link" onClick={handleFairUseClick}>
                Site's Fair Use Policy
            </span>
            {showModal && (
                <div className="modal">
                    <div className="modal-content">
                        <span className="close" onClick={handleCloseModal}>
                            &times;
                        </span>
                        <h2>Fair Use Policy</h2>
                        <img src={Gastly} alt="Gastly" className="image" />
                        <br></br>
                        <p>

                            Our site uses trademarks and other intellectual property that belong to their respective owners. However, our use of these trademarks is protected under fair dealing provisions in the Copyright, Designs and Patents Act 1988, as we are a non-profit and open source project.

                            We acknowledge and respect the intellectual property rights of the owners of the trademarks and other intellectual property that we use. We use the trademarks and other intellectual property only to the extent necessary to identify the products or services to which they refer. We do not use the trademarks in a way that would suggest an endorsement or affiliation with the owners of the trademarks.

                            Our use of the trademarks and other intellectual property is non-commercial in nature, and we do not make any profit from our use of them. Our use of the trademarks and other intellectual property is also transformative in nature, as we are using them to create something new and different from their original purpose.

                            If you are the owner of any of the trademarks or other intellectual property used on our site and you believe that our use of them does not fall under fair dealing provisions or otherwise infringes your intellectual property rights, please contact us at johgo.search@gmail.com. We will promptly review your request and take appropriate action.

                            By using our site, you agree to comply with our fair use policy and respect the intellectual property rights of others.</p>
                    </div>
                </div>
            )}
        </div>
    );
}

export default Footer;
import React, { useState, useEffect } from 'react';
import '../../css/Footer.css';
import Gastly from '../../images/gastly.png'
import { Button, Modal, ConfigProvider, message } from 'antd';
import { GithubOutlined, CopyOutlined } from '@ant-design/icons';
import { CopyToClipboard } from "react-copy-to-clipboard";

function Footer() {
    const [isModalOpen, setIsModalOpen] = useState(false);
    const [isFairUseModalOpen, setIsFairUseModalOpen] = useState(false);
    const [messageApi, contextHolder] = message.useMessage();
    const [opacity, setOpacity] = useState(1);

    useEffect(() => {
        const handleScroll = () => {
            const newOpacity = window.scrollY > 100 ? Math.max(1 - (window.scrollY - 100) / 300, 0) : 1;
            setOpacity(newOpacity);
        };

        window.addEventListener('scroll', handleScroll);

        return () => {
            window.removeEventListener('scroll', handleScroll);
        };
    }, []);

    const success = () => {
        messageApi.open({
            type: 'success',
            content: 'Email copied to clipboard!'
        });
    };
    const showModal = () => {
        setIsModalOpen(true);
    };
    const handleOk = () => {
        setIsModalOpen(false);
    };
    const handleCancel = () => {
        setIsModalOpen(false);
    };
    const showFUModal = () => {
        setIsFairUseModalOpen(true);
    };
    const handleFUCancel = () => {
        setIsFairUseModalOpen(false);
    };

    const copyToClipboard = () => {
        setIsModalOpen(false);
        success();
    }

    return (
        <div className="footer">
            {contextHolder}
            <ConfigProvider
                theme={{
                    token: {
                        colorPrimary: '#222',
                        colorBgContainer: '#222',
                        colorBgElevated: '#222',
                        colorTextBase: 'white',
                        borderRadius: 20,
                    },
                }}>
                <span onClick={showFUModal} className="footer-link" style={{ opacity: opacity }}>
                    Site's Fair Use Policy
                </span>
                <Modal
                    title={<h2>Sites Fair Use Policy <img src={Gastly} alt="Gastly" className="image" /></h2>}
                    open={isFairUseModalOpen}
                    centered
                    onCancel={handleFUCancel}
                    footer={[

                    ]}
                >
                    <p>Our site uses trademarks and other intellectual property that belong to their respective owners. However, our use of these trademarks is protected under fair dealing provisions in the Copyright, Designs and Patents Act 1988, as we are a non-profit and open source project.

                        We acknowledge and respect the intellectual property rights of the owners of the trademarks and other intellectual property that we use. We use the trademarks and other intellectual property only to the extent necessary to identify the products or services to which they refer. We do not use the trademarks in a way that would suggest an endorsement or affiliation with the owners of the trademarks.

                        Our use of the trademarks and other intellectual property is non-commercial in nature, and we do not make any profit from our use of them. Our use of the trademarks and other intellectual property is also transformative in nature, as we are using them to create something new and different from their original purpose.

                        If you are the owner of any of the trademarks or other intellectual property used on our site and you believe that our use of them does not fall under fair dealing provisions or otherwise infringes your intellectual property rights, please contact us at johgo.search@gmail.com. We will promptly review your request and take appropriate action.</p>
                </Modal>
                <span onClick={showModal} className="footer-link" style={{ opacity: opacity }}>
                    About JohGo
                </span>
                <Modal
                    title={<h2>About JohGo</h2>}
                    open={isModalOpen}
                    centered
                    onCancel={handleCancel}
                    footer={[
                        <CopyToClipboard text="johgo.search@gmail.com"
                            onCopy={() => copyToClipboard} >
                            <Button
                                onClick={copyToClipboard}
                                type="primary"
                            >Copy email<CopyOutlined />
                            </Button>
                        </CopyToClipboard>,
                        <Button
                            key="link"
                            href="https://github.com/tomronw/johgo"
                            target='_blank'
                            type="primary"
                            onClick={handleOk}
                        >JohGo Github<GithubOutlined />
                        </Button>,
                    ]}>
                    <p>No ads, no affiliate links, no mystery boxes. This is a completely free, open-source and community driven project for searching for in stock TCG products in the UK. </p>
                    <p>See a bug or error? Please report it via email or raising an issue on Github via the button below! </p>
                </Modal>
            </ConfigProvider >
        </div >
    );
}

export default Footer;
import HomeSearchBox from "../components/home/HomeSearchBox.jsx";
import SearchContext from "../context/SearchContext.js";
import Loading from "../components/miscellaneous/Loading.jsx";
import { useContext, useState } from "react";
import '../css/App.css';
import ShowResults from "../components/results/ShowResults.jsx";
import ReleaseCalendar from "../components/home/ReleaseCalendar.jsx";
import { Popover, Button, Tooltip, ConfigProvider, theme } from 'antd';
import { useTheme } from '../context/ThemeContext';
import { CalendarTwoTone } from '@ant-design/icons';

function Home() {
    const { products, loading } = useContext(SearchContext);
    const [isCalendarVisible, setIsCalendarVisible] = useState(false);

    const { darkMode } = useTheme();
    var currentTheme

    if (darkMode) {
        currentTheme = theme.darkAlgorithm
    } else {
        currentTheme = theme.lightAlgorithm
    }

    if (!loading) {
        return (
            <ConfigProvider
                theme={{
                    algorithm: currentTheme,
                }}
            >
                <div>
                    <div className="App">
                        {products.length > 0 ? <ShowResults /> :
                            <>
                                <HomeSearchBox />
                                <div style={{ textAlign: 'center', marginTop: '-7%', animation: 'slideFromBottom 1s forwards' }}>
                                    <Popover
                                        content={<ReleaseCalendar />}
                                        title="Pokémon TCG Product Release Calendar"
                                        trigger="click"
                                        visible={isCalendarVisible}
                                        onVisibleChange={setIsCalendarVisible}
                                    >
                                        <Tooltip title="TCG Product Release Calendar">
                                            <Button size='large' onClick={() => setIsCalendarVisible(!isCalendarVisible)} icon={<CalendarTwoTone />} style={{ opacity: '0.7' }} />
                                        </Tooltip>
                                    </Popover>
                                </div>
                            </>
                        }
                    </div>
                </div>
            </ConfigProvider>
        )
    } else {
        return (
            <div className="App" style={{ position: 'relative' }}>
                <Loading />
            </div>
        );
    }
}

export default Home;

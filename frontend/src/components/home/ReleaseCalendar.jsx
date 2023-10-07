import React, { useState, useEffect } from 'react';
import dayjs from 'dayjs';
import 'dayjs/locale/en-gb';
import dayLocaleData from 'dayjs/plugin/localeData';
import { Calendar, Popover, Badge, theme, ConfigProvider } from 'antd';
import { useTheme } from '../../context/ThemeContext';
import CalendarService from '../../services/CalendarService';
dayjs.extend(dayLocaleData);

export default function ReleaseCalendar() {
    const [visiblePopoverDate, setVisiblePopoverDate] = useState(null);
    const { darkMode } = useTheme();
    var currentTheme

    if (darkMode) {
        currentTheme = theme.darkAlgorithm
    } else {
        currentTheme = theme.lightAlgorithm
    }

    const [calendarData, setCalendarData] = useState([]);
    // eslint-disable-next-line
    const [error, setError] = useState(null);

    useEffect(() => {
        CalendarService.getCalendar()
            .then(data => {
                setCalendarData(data);
            })
            .catch(err => {
                console.error(err);
                setError(err);
            });
    }, []);

    const specialDates = calendarData.reduce((acc, item) => {
        acc[item.date] = item.releases;  // changed from item.release to item.releases
        return acc;
    }, {});

    const dateCellRender = (value) => {
        const dateStr = value.format('YYYY-MM-DD');
        const isSpecialDate = !!specialDates[dateStr];
        const isToday = dayjs().format('YYYY-MM-DD') === dateStr; // Check if the current date is today
        const popoverContent = specialDates[dateStr]?.map(release => (
            <div key={release}>
                • {release}
            </div>
        ));

        const baseStyle = {
            textAlign: 'center',
            position: 'relative',
            height: '100%',
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'center',
            borderRadius: '7px',
        };

        if (isToday) {
            return (
                <div
                    style={{ ...baseStyle, backgroundColor: '#2825e6', fontWeight: 'bold' }}
                >
                    <span>{value.date()}</span>
                </div>
            );
        } else if (isSpecialDate) {
            return (
                <div
                    style={{ ...baseStyle, backgroundColor: '#f28d92' }}
                    onClick={() => setVisiblePopoverDate(dateStr)}
                >
                    <Popover
                        content={popoverContent}
                        title={value.format('MMMM D, YYYY')}
                        trigger="click"
                        open={visiblePopoverDate === dateStr}
                        onOpenChange={visible => !visible && setVisiblePopoverDate(null)}
                    >
                        <span>{value.date()}</span>
                        <Badge dot style={{ position: 'absolute', right: 20, verticalAlign: 'middle' }} />
                    </Popover>
                </div>
            );
        } else {
            return (
                <div style={baseStyle}>
                    <span>{value.date()}</span>
                </div>
            );
        }
    };

    return (
        <ConfigProvider
            theme={{
                algorithm: currentTheme,
            }}
        >
            <div style={{ maxWidth: 300 }}>

                <Calendar
                    fullscreen={false}
                    fullCellRender={dateCellRender}
                />
            </div>
        </ConfigProvider>
    );
}

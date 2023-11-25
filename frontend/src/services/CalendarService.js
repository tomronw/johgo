const defaultHeaders = {
    'Content-Type': 'application/json',
    'Accept': 'application/json',
};

class CalendarService {
    static getCalendar() {
        const options = {
            method: 'GET',
            headers: defaultHeaders,
        };

        return fetch(`${process.env.REACT_APP_BASEAPI}/serve_calendar`, options)
            .then(response => {
                if (!response.ok) {
                    throw new Error(`Error: ${response.statusText}`);
                }
                return response.json();
            });
    }
}

export default CalendarService;

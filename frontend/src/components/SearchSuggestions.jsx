import { useState, useEffect } from 'react';


export function useFetchSearchSuggestions() {
    const [searchSuggestions, setSearchSuggestions] = useState([]);

    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await fetch(process.env.REACT_APP_SUGGESTIONS);
                if (!response.ok) {
                    throw new Error(`HTTP error! status: ${response.status}`);
                }
                const data = await response.json();
                setSearchSuggestions(data);
            } catch (error) {
                console.error('Error fetching data: ', error);
            }
        };

        fetchData();
    }, []);

    return searchSuggestions;
}
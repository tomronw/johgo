import { createContext, useReducer } from 'react';
import searchReducer from "./SearchReducer";

const SearchContext = createContext();

export const SearchProvider = ({ children }) => {

    const initialState = {
        products: [],
        loading: false
    }


    const [state, dispatch] = useReducer(searchReducer, initialState)


    const fetchProducts = async (query, checked) => {
        clearProducts()
        setLoading()
        const APIENDPOINT = process.env.REACT_APP_APIENDPOINT;

        try {
            const response = await fetch(`${APIENDPOINT}/v1/search?query=${query.toString()}&filter_singles=${checked.toString()}`, {
                headers: {
                    "content-type": "application/json"
                }
            })

            if (response.status === 200) {
                const data = await response.json()
                dispatch({ type: "FETCH_PRODUCTS", payload: data.Data.products })
            } else {
                dispatch({ type: "FETCH_PRODUCTS", payload: [{}] })
            }
        } catch (error) {
            dispatch({ type: "FETCH_PRODUCTS", payload: [{}] })
        }
    }

    // clear products
    const clearProducts = () => {
        dispatch({ type: "CLEAR_PRODUCTS" })
    }

    const setLoading = () => {
        dispatch({ type: "SET_LOADING" })
    }

    return (<SearchContext.Provider value={{ products: state.products, loading: state.loading, fetchProducts, clearProducts }}>{children}</SearchContext.Provider>)
}

export default SearchContext;
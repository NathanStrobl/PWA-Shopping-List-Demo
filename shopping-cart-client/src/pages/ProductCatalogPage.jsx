import { useEffect, useState } from "react";
import ItemRow from "../components/ItemRow";

export default function ProductCatalogPage() {
    const apiUrl = import.meta.env.VITE_API_URL;
    const queryParams = new URLSearchParams(window.location.search);

    const [categoryProducts, setCategoryProducts] = useState([]);
    const [displayMode, setDisplayMode] = useState(0);
    // 0 - loading
    // 1 - normal
    // 2 - failed

    useEffect(() => {
        fetch(`${apiUrl}?category=${queryParams.get("category")}`)
        .then(response => {
            if(response.status == 200) return response.json();
        }).then(data => {
            setCategoryProducts(data.products);
            setDisplayMode(1);
        }).catch(() => setDisplayMode(2));
    }, []);

    return (
        <>
            <div style={{display: displayMode == 0 ? 'block' : 'none'}}>
                <h2>Loading products...</h2>
            </div>
            <div className="items-container" style={{display: displayMode == 1 ? 'block' : 'none'}}>
                {categoryProducts.map(item => (
                    <ItemRow itemObj={item} />
                ))}
            </div>
            <div style={{display: displayMode == 2 ? 'block' : 'none'}}>
                <h2>Loading failed!</h2>
                <h3>Walmark's is sorry for the inconvenience, please try again later.</h3>
            </div>
        </>
    );
}
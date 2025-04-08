import { useEffect, useState } from "react";
import CartRow from "../components/CartRow";
import Cookies from 'js-cookie';

export default function CartPage() {
    const [cartContents, setCartContents] = useState([]);

    useEffect(() => {
        updateList();
    }, [])

    function updateList() {
        const cartContentsRaw = Cookies.get('walmarks-cart');
        if(cartContentsRaw) {
            setCartContents(JSON.parse(cartContentsRaw));
        }
    }

    return (
        <>
            <div style={{display: cartContents.length == 0 ? 'block' : 'none'}}>
                <h2>No items in cart.</h2>
                <h3>Go back to the homepage to start adding products to your cart.</h3>
            </div>
            <div className="items-container" style={{display: cartContents.length > 0 ? 'block' : 'none'}}>
                {cartContents.map(item => (
                    <CartRow cartObj={item} updateCallback={updateList} />
                ))}
            </div>
        </>
    );
}
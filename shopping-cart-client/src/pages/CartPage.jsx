import CartRow from "../components/CartRow";

export default function CartPage() {
    const sampleCartObj = {
        title: "Test",
        price: 1.99,
        aisleNo: "A23",
        upc: "1234567890"
    };

    const sampleCartObj2 = {
        title: "iPhone 14 Pro Max (128GB)",
        price: 1199.99,
        aisleNo: 'A23',
        upc: '1234567890'
    }

    return (
        <div className="items-container">
            <CartRow cartObj={sampleCartObj} />
            <CartRow cartObj={sampleCartObj2} />
        </div>
    );
}
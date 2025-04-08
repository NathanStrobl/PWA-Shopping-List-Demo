import ItemRow from "../components/ItemRow";

export default function ProductCatalogPage() {
    const sampleItemObj = {
        title: "iPhone 14 Pro Max (128GB)",
        price: 1199.99,
        quantity: 50,
        aisleNo: 'A23',
        upc: '1234567890'
    }

    const sampleItemObj2 = {
        title: "Test",
        price: 1.99,
        quantity: 50,
        aisleNo: 'A23',
        upc: '1234567890'
    }

    return (
        <div className="items-container">
            <ItemRow itemObj={sampleItemObj} />
            <ItemRow itemObj={sampleItemObj2} />
        </div>
    );
}
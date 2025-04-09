import CategoryButton from "../components/CategoryButton";
import AppleImage from '/images/apple.jpg';
import GoogleHomeImage from '/images/googlehome.png';
import PencilsImage from '/images/pencils.jpg';
import ShirtImage from '/images/shirt.png';
import SlinkyImage from '/images/slinky.jpg';
import SofaImage from '/images/sofa.jpg';

export default function Homepage() {
    return (
        <>
            <div className="welcome-message-container">
                <h1 class="shrink" style={{lineHeight: '20px'}}>Welcome to Walmark's!</h1>
                <h3>Select one of the categories below to begin browsing our selection.</h3>
            </div>
            <div className="categories-container">
                <CategoryButton title="Clothing" image={ShirtImage} path='/browse?category=Clothing' />
                <CategoryButton title="Home Goods" image={SofaImage} path='browse?category=Home%20Goods' />
                <CategoryButton title="Electronics" image={GoogleHomeImage} path='browse?category=Electronics' />
                <CategoryButton title="Office Supplies" image={PencilsImage} path='browse?category=Office%20Supplies' />
                <CategoryButton title="Groceries" image={AppleImage} path='browse?category=Groceries' />
                <CategoryButton title="Toys" image={SlinkyImage} path='browse?category=Toys' />
            </div>
        </>
    );
}
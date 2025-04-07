import CategoryButton from "../components/CategoryButton";
import WalmarksLogo from '../assets/walmarks.png';

export default function Homepage() {
    return (
        <>
            <div className="welcome-message-container">
                <h1 style={{lineHeight: '20px'}}>Welcome to Walmark's!</h1>
                <h3>Select one of the categories below to begin browsing our selection.</h3>
            </div>
            <div className="categories-container">
                <CategoryButton title="Clothing" image={WalmarksLogo} />
                <CategoryButton title="Home Goods" image={WalmarksLogo} />
                <CategoryButton title="Electronics" image={WalmarksLogo} />
                <CategoryButton title="Office Supplies" image={WalmarksLogo} />
                <CategoryButton title="Groceries" image={WalmarksLogo} />
                <CategoryButton title="Toys" image={WalmarksLogo} />
                <CategoryButton title="Health" image={WalmarksLogo} />
                <CategoryButton title="Hygiene" image={WalmarksLogo} />
                <CategoryButton title="Automotive" image={WalmarksLogo} />
            </div>
        </>
    );
}
export default function CategoryButton({ title, image, path }) {
    return (
        <a href={path} className="category-button">
            <img src={image} />
            <h2>{title}</h2>
        </a>
    );
}
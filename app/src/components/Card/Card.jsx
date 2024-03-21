import './card.scss';

const Card = ({ title, description }) => {
  return (
    <div className='card'>
      <div>
        <p className='card-title'>{title}</p>
      </div>
      <h2 className='card-description'>{description}</h2>
    </div>
  );
};

export default Card;

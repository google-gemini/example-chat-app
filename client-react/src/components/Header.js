import React from 'react';

const Header = ({ toggled, setToggled }) => {
  return (
    <div className="chat-header">
      <h1>Example chat app</h1>
      <span className='toggle-text'>Stream Response</span>
      <button 
        className={`toggle-btn ${toggled ? "toggled": ""}`}
        onClick={() => setToggled(!toggled)}
      >
        <div className="toggle-hover">
          <div className='thumb'></div>
          {toggled === false ? (
            <span className="toggle-hover-text">Streaming response Off</span>
          ) : (
            <span className="toggle-hover-text">Streaming response On</span>
          )}
        </div>
      </button>
    </div>
  );
};

export default Header;
import React from 'react';
import PropTypes from 'prop-types';

const ListView = ({ names, title }) => (
  <div>
    <p>{title}</p>
    <ul>{names.map((name, index) => <li key={index}>{name}</li>)}</ul>
  </div>
);

ListView.propsTypes = {
  names: PropTypes.arrayOf(PropTypes.string),
  title: PropTypes.string,
};

export default ListView;

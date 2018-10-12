import React, { Component } from 'react';
import PropTypes from 'prop-types';

class ListView extends Component {
  componentDidMount() {
    this.props.onMount();
  }

  render() {
    const { names, title } = this.props;
    return (
      <div>
        <p>{title}</p>
        <ul>{names.map((name, index) => <li key={index}>{name}</li>)}</ul>
      </div>
    );
  }
}

ListView.propsTypes = {
  names: PropTypes.arrayOf(PropTypes.string),
  onMount: PropTypes.func,
  title: PropTypes.string,
};

ListView.defaultProps = {
  onMount: () => {},
};

export default ListView;

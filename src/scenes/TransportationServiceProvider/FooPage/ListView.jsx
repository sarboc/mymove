import React, { Component } from 'react';
import PropTypes from 'prop-types';

class ListView extends Component {
  componentDidMount() {
    this.props.onMount();
  }

  render() {
    const { loading, names, title } = this.props;
    return (
      <div>
        <p>{title}</p>
        {loading && <div>Loading...</div>}
        {!loading && (
          <ul>{names.map((name, index) => <li key={index}>{name}</li>)}</ul>
        )}
      </div>
    );
  }
}

ListView.propsTypes = {
  loading: PropTypes.bool,
  names: PropTypes.arrayOf(PropTypes.string),
  onMount: PropTypes.func,
  title: PropTypes.string,
};

ListView.defaultProps = {
  onMount: () => {},
};

export default ListView;

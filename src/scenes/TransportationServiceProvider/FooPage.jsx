import React from 'react';
import { connect } from 'react-redux';

const FooPage = ({ id }) => <div>Hello!!!! {id}</div>;

const mapStateToProps = (_state, ownProps) => ({
  id: ownProps.match.params.id,
});

export default connect(mapStateToProps)(FooPage);

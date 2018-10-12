import React from 'react';
import { connect } from 'react-redux';
import ListView from './FooPage/ListView';

const FooPage = ({ id }) => (
  <div style={{ paddingLeft: '20px' }}>
    <div>Hello!!!! {id}</div>
    <ListView
      names={['Rebecca', 'Chris', 'Ron', 'Kara', 'Sara']}
      title="Team Roci"
    />
    <ListView names={['Alexi', 'Jim', 'Kim', 'Reggie']} title="A-Team" />
    <ListView
      names={['Andrea', 'Donald', 'Erin', 'Patrick']}
      title="Team Teen Vogue"
    />
  </div>
);

const mapStateToProps = (_state, ownProps) => ({
  id: ownProps.match.params.id,
});

export default connect(mapStateToProps)(FooPage);

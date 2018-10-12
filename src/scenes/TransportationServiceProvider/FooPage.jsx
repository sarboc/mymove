import React from 'react';
import { connect } from 'react-redux';
import ATeam from './FooPage/ATeamContainer';
import Roci from './FooPage/TeamRociContainer';
import TeenVogue from './FooPage/TeenVogueContainer';

const FooPage = ({ aTeam, id, roci, teenVogue }) => (
  <div style={{ paddingLeft: '20px' }}>
    <div>Hello!!!! {id}</div>
    <ATeam id={id} />
    <Roci id={id} />
    <TeenVogue id={id} />
  </div>
);

const mapStateToProps = (state, ownProps) => ({
  id: ownProps.match.params.id,
});

export default connect(mapStateToProps)(FooPage);

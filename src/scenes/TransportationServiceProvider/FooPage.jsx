import React from 'react';
import { connect } from 'react-redux';
import ListView from './FooPage/ListView';

const FooPage = ({ aTeam, id, roci, teenVogue }) => (
  <div style={{ paddingLeft: '20px' }}>
    <div>Hello!!!! {id}</div>
    <ListView names={roci.members} title={roci.name} />
    <ListView names={aTeam.members} title={aTeam.name} />
    <ListView names={teenVogue.members} title={teenVogue.name} />
  </div>
);

const mapStateToProps = (state, ownProps) => ({
  id: ownProps.match.params.id,
  aTeam: state.tsp.teams.a_team,
  roci: state.tsp.teams.roci,
  teenVogue: state.tsp.teams.teen_vogue,
});

export default connect(mapStateToProps)(FooPage);

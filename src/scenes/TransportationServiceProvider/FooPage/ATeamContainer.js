import { connect } from 'react-redux';
import ListView from './ListView';

const mapStateToProps = state => {
  const team = state.tsp.teams.a_team;
  return {
    names: team.members,
    title: team.name,
  };
};

export default connect(mapStateToProps)(ListView);

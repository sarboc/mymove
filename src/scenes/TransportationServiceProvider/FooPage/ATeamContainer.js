import { connect } from 'react-redux';
import ListView from './ListView';
import { loadTeam } from '../ducks';

const mapStateToProps = state => {
  const team = state.tsp.teams.a_team;
  return {
    loading: team.loading,
    names: team.members,
    title: team.name,
  };
};

const mapDispatchToProps = (dispatch, { id }) => ({
  onMount: () => setTimeout(() => dispatch(loadTeam('a_team')), 4000),
});

export default connect(mapStateToProps, mapDispatchToProps)(ListView);

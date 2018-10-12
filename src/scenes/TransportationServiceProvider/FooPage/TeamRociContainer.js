import { connect } from 'react-redux';
import ListView from './ListView';

const mapStateToProps = state => {
  const team = state.tsp.teams.roci;
  return {
    names: team.members,
    title: team.name,
  };
};

const mapDispatchToProps = (_dispatch, { id }) => ({
  onMount: () => alert(`Whaaaaaaat?! ${id}`),
});

export default connect(mapStateToProps, mapDispatchToProps)(ListView);

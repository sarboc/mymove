import { connect } from 'react-redux';
import ListView from './ListView';

const mapStateToProps = state => {
  const team = state.tsp.teams.a_team;
  return {
    names: team.members,
    title: team.name,
  };
};

const mapDispatchToProps = (_dispatch, { id }) => ({
  onMount: () => console.log('Hellloooooooo ', id),
});

export default connect(mapStateToProps, mapDispatchToProps)(ListView);

import { connect } from 'react-redux';
import ListView from './ListView';

const mapStateToProps = state => {
  const team = state.tsp.teams.teen_vogue;
  return {
    names: team.members,
    title: team.name,
  };
};

export default connect(mapStateToProps)(ListView);

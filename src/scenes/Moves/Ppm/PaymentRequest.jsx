import React, { Component, Fragment } from 'react';
import PropTypes from 'prop-types';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import Alert from 'shared/Alert'; // eslint-disable-line
import { get } from 'lodash';
import { includes } from 'lodash';
import qs from 'query-string';

import DocumentUploader from 'shared/DocumentViewer/DocumentUploader';
import { convertDollarsToCents } from 'shared/utils';
import { createMoveDocument } from 'shared/Entities/modules/moveDocuments';
import { createMovingExpenseDocument } from 'shared/Entities/modules/movingExpenseDocuments';

import {
  selectAllDocumentsForMove,
  getMoveDocumentsForMove,
} from 'shared/Entities/modules/moveDocuments';

import { submitExpenseDocs } from './ducks.js';

import './PaymentRequest.css';

function RequestPaymentSection(props) {
  const { ppm, updatingPPM, submitDocs, disableSubmit } = props;

  if (!ppm) {
    return null;
  }

  if (ppm.status === 'APPROVED') {
    return (
      <Fragment>
        <h4>Done uploading documents?</h4>
        <button
          onClick={submitDocs}
          className="usa-button"
          disabled={updatingPPM || disableSubmit}
        >
          Submit Payment Request
        </button>
      </Fragment>
    );
  } else if (ppm.status === 'PAYMENT_REQUESTED') {
    return (
      <Fragment>
        <h4>Payment requested, awaiting approval.</h4>
      </Fragment>
    );
  } else {
    console.error(
      'Unexpectedly got to PaymentRequest screen without PPM approval',
    );
  }
}

export class PaymentRequest extends Component {
  constructor(props) {
    super(props);
    this.submitDocs = this.submitDocs.bind(this);
  }

  componentDidMount() {
    this.props.getMoveDocumentsForMove(this.props.match.params.moveId);
  }

  submitDocs() {
    this.props
      .submitExpenseDocs()
      .then(() => {
        this.props.history.push('/');
      })
      .catch(() => {
        window.scrollTo(0, 0);
      });
  }

  handleSubmit = (uploadIds, formValues) => {
    const { currentPpm } = this.props;
    if (get(formValues, 'move_document_type', false) === 'EXPENSE') {
      formValues.requested_amount_cents = convertDollarsToCents(
        formValues.requested_amount_cents,
      );
      return this.props.createMovingExpenseDocument(
        this.props.match.params.moveId,
        currentPpm.id,
        uploadIds,
        formValues.title,
        formValues.moving_expense_type,
        formValues.move_document_type,
        formValues.requested_amount_cents,
        formValues.payment_method,
        formValues.notes,
      );
    }
    return this.props.createMoveDocument(
      this.props.match.params.moveId,
      currentPpm.id,
      uploadIds,
      formValues.title,
      formValues.move_document_type,
      formValues.notes,
    );
  };

  render() {
    const { location, moveDocuments, updateError, docTypes } = this.props;
    const numMoveDocs = get(moveDocuments, 'length', 'TBD');
    const disableSubmit = numMoveDocs === 0;
    const moveDocumentType = qs.parse(location.search).moveDocumentType;
    const initialValues = {};

    // Verify the provided doc type against the schema
    if (includes(docTypes, moveDocumentType)) {
      initialValues.move_document_type = moveDocumentType;
    }

    return (
      <div className="usa-grid payment-request">
        <div className="usa-width-two-thirds">
          {updateError && (
            <div className="usa-width-one-whole error-message">
              <Alert type="error" heading="An error occurred">
                There was an error requesting payment, please try again.
              </Alert>
            </div>
          )}
          <h2>Request Payment </h2>
          <div className="instructions">
            Please upload all your weight tickets, expenses, and storage fee
            documents one at a time. For expenses, you’ll need to enter
            additional details.
          </div>
          <DocumentUploader
            form="payment-docs"
            genericMoveDocSchema={this.props.genericMoveDocSchema}
            initialValues={initialValues}
            isPublic={false}
            location={location}
            moveDocSchema={this.props.moveDocSchema}
            onSubmit={this.handleSubmit}
          />
          <RequestPaymentSection
            ppm={this.props.currentPpm}
            updatingPPM={this.props.updatingPPM}
            submitDocs={this.submitDocs}
            disableSubmit={disableSubmit}
          />
        </div>
        <div className="usa-width-one-third">
          <h4 className="doc-list-title">All Documents ({numMoveDocs})</h4>
          {(moveDocuments || []).map(doc => {
            return (
              <div className="panel-field" key={doc.id}>
                <span>{doc.title}</span>
              </div>
            );
          })}
        </div>
      </div>
    );
  }
}
PaymentRequest.propTypes = {
  docTypes: PropTypes.arrayOf(PropTypes.string),
  moveDocuments: PropTypes.arrayOf(PropTypes.object),
  genericMoveDocSchema: PropTypes.object.isRequired,
  moveDocSchema: PropTypes.object.isRequired,
};

const mapStateToProps = (state, props) => ({
  moveDocuments: selectAllDocumentsForMove(state, props.match.params.moveId),
  currentPpm: state.ppm.currentPpm,
  updatingPPM: state.ppm.hasSubmitInProgress,
  updateError: state.ppm.hasSubmitError,
  docTypes: get(
    state,
    'swaggerInternal.spec.definitions.MoveDocumentType.enum',
    [],
  ),
  genericMoveDocSchema: get(
    state,
    'swaggerInternal.spec.definitions.CreateGenericMoveDocumentPayload',
    {},
  ),
  moveDocSchema: get(
    state,
    'swaggerInternal.spec.definitions.MoveDocumentPayload',
    {},
  ),
});

const mapDispatchToProps = dispatch =>
  bindActionCreators(
    {
      getMoveDocumentsForMove,
      submitExpenseDocs,
      createMoveDocument,
      createMovingExpenseDocument,
    },
    dispatch,
  );

export default connect(mapStateToProps, mapDispatchToProps)(PaymentRequest);

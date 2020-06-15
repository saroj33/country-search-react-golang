import React, { PureComponent } from "react";
import PropTypes from "prop-types";
import "../css/ResultRow.css";
import ReactCountryFlag from "react-country-flag"

export default class ResultsRow extends PureComponent {
  static propTypes = {
    name: PropTypes.string,
    code: PropTypes.string,
    matching: PropTypes.array
  };

  render() {
      return (
      <div className="component-country-result-row">
        <ReactCountryFlag countryCode={this.props.code} />
        <span className="title">{this.props.name}</span>
        {this.props.matching && this.props.matching.map((data) => (
        <span className="matches">{data.key}:{data.val}</span>
    ))}
      </div>
    );
  }
}

import React, { PureComponent } from "react";
import PropTypes from "prop-types";


import ResultRow from "./ResultRow";
import "../css/Results.css";

export default class Results extends PureComponent {
  static propTypes = {
    countryData: PropTypes.array
  };

  render() {
    return (
      <div className="component-country-results">
        {this.props.countryData && this.props.countryData.map(countryData => (
          <ResultRow
            name={countryData.name}
            code={countryData.code}
            matching={countryData.match}
          />
        ))}
      </div>
    );
  }
}

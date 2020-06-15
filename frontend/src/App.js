import React, { PureComponent } from "react";
import Header from "./components/Header";
import SearchInput from "./components/SearchInput";
import Results from "./components/Results";

export default class App extends PureComponent {
  constructor(props) {
    super(props);
    
    this.state = {
      filteredCountry:[]
    };
  }
  componentDidMount() {
     this.filterCountry("")
  }
  filterCountry(searchText){
  fetch("http://localhost:2020/search?text="+searchText)
  .then(res => res.json())
  .then(json => this.setState({ filteredCountry: json }));
}
  handleSearchChange = event => {
    this.filterCountry(event.target.value)
  };

  render() {
    return (
      <div>
        <Header />
        <SearchInput textChange={this.handleSearchChange} />
        <Results countryData={this.state.filteredCountry} />
      </div>
    );
  }
}

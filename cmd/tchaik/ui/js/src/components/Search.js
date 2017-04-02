"use strict";

import React from "react";

import SearchStore from "../stores/SearchStore.js";

import {GroupList as GroupList} from "./Collection.js";

import Icon from "./Icon.js";

const Search = () => {
  return (
    <div className="collection">
      <Results />
    </div>
  );
}

export default Search;


function getResultsState() {
  return {results: SearchStore.getResults()};
}

class Results extends React.Component {
  constructor(props) {
    super(props);

    this.state = getResultsState();
    this._onChange = this._onChange.bind(this);
  }

  componentDidMount() {
    SearchStore.addChangeListener(this._onChange);
  }

  componentWillUnmount() {
    SearchStore.removeChangeListener(this._onChange);
  }

  render() {
    var list = this.state.results;
    if (list.length === 0) {
      return <div className="no-results"><Icon icon="audiotrack" />No results found</div>;
    }
    return <GroupList path={["Root"]} list={list} depth={0} />;
  }

  _onChange() {
    this.setState(getResultsState());
  }
}

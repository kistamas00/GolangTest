import * as React from 'react';
import {Provider} from "react-redux";
import {Route, Switch} from 'react-router';
import AdminPage from "./components/pages/AdminPage";
import PublicPage from "./components/pages/PublicPage";
import Store from './store/Store';

class App extends React.Component {
  public render() {
    return (
      <Provider store={Store}>
        <Switch>
          <Route exact={true} path='/admin' component={AdminPage} />
          <Route exact={true} path='/letsVote/:id' component={PublicPage} />
        </Switch>
      </Provider>
    );
  }
}

export default App;

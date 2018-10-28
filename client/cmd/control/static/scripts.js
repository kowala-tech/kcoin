class MintListEntry extends React.Component {
  constructor(props) {
    super(props);
    this.onConfirm = this.onConfirm.bind(this);
  }
  render() {
    return (
      <tr>
        <td>{this.props.entry.id}</td>
        <td>{this.props.entry.to}</td>
        <td>{(this.props.entry.amount / 1.0e18).toFixed(2)} ({this.props.entry.amount})</td>
        <td>{this.props.entry.confirmed ? "Confirmed" : "Pending"}</td>
        <td>
          {!this.props.entry.confirmed &&
            <a href="#" class="button" onClick={this.onConfirm}>Confirm</a>
          }
        </td>
        
      </tr>
    );
  }
  onConfirm(e) {
    e.preventDefault()
    this.props.sendData({
      action: "confirm_mint",
      governor: this.props.governor,
      id: this.props.entry.id,
    })
  }
}

class MintList extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      idSearch: "",
      addrSearch: "",
      statusSearch: "0",
    }
    this.handleIdSearchChange = this.handleIdSearchChange.bind(this);
    this.handleAddrSearchChange = this.handleAddrSearchChange.bind(this);
    this.handleStatusSearchChange = this.handleStatusSearchChange.bind(this);
  }

  entries() {
    return this.props.entries.filter( (entry) => {
      if (this.state.idSearch != "" && entry.id != this.state.idSearch) {
        return false
      }
      if (this.state.addrSearch != "" && entry.to.indexOf(this.state.addrSearch) == -1) {
        return false
      }
      if (this.state.statusSearch == "1" && !entry.confirmed) {
        return false
      }
      if (this.state.statusSearch == "2" && entry.confirmed) {
        return false
      }
      return true
    })
  }

  handleIdSearchChange(e) {
    this.setState({idSearch: e.target.value})
  }

  handleAddrSearchChange(e) {
    this.setState({addrSearch: e.target.value})
  }

  handleStatusSearchChange(e) {
    this.setState({statusSearch: e.target.value})
  }

  render() {
    return (
      <table>
        <thead>
          <tr>
            <th>ID <input placeholder="all" type="text" value={this.state.idSearch} onChange={this.handleIdSearchChange} style={{width: 40}}/></th>
            <th>Address <input placeholder="all" type="text" value={this.state.addrSearch} onChange={this.handleAddrSearchChange}/></th>
            <th>Amount</th>
            <th>Status <select value={this.state.statusSearch} onChange={this.handleStatusSearchChange}>
              <option value="0">All</option>
              <option value="1">Confirmed</option>
              <option value="2">Pending</option>
            </select></th>
            <th>Actions</th>
          </tr>
        </thead>
        <tbody>
          {this.entries().map((entry) => (
            <MintListEntry key={entry.id} entry={entry} governor={this.props.governor} sendData={this.props.sendData} />
          ))}
        </tbody>
      </table>
    )
  }
}

class Minting extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      governor: "",
      address: "",
      amount: "",
      unit: "18",
    }
    this.handleGovernorChange = this.handleGovernorChange.bind(this);
    this.handleAddressChange = this.handleAddressChange.bind(this);
    this.handleAmountChange = this.handleAmountChange.bind(this);
    this.handleUnitChange = this.handleUnitChange.bind(this);
    this.handleProposeMint = this.handleProposeMint.bind(this);
  }

  componentWillReceiveProps(nextProps) {
    if (this.state.governor == "" && nextProps.accounts && nextProps.accounts.length > 0) {
      this.setState({ governor: nextProps.accounts[0] });
    }
  }

  render() {
    return (
      <fieldset>
        <h3>Mint tokens</h3>
        <label for="governor_account">Governor account</label>

        <select type="text" onChange={this.handleGovernorChange} value={this.state.governor}>
          {(this.props.accounts||[]).map((account) => (
            <option key={account} value={account}>{account}</option>
          ))}
        </select>

        <form onSubmit={this.handleProposeMint}>
          <label for="mint_address">Recipient ddress</label>
          <input type="text" id="mint_address" onChange={this.handleAddressChange} value={this.state.address}/>
          <label for="mint_amount">Amount</label>
          <input type="text" id="mint_amount" onChange={this.handleAmountChange} value={this.state.amount}/>
          <label for="mint_unit">Unit</label>
          <select type="text" id="mint_unit" onChange={this.handleUnitChange} value={this.state.unit}>
            <option value="18">mToken</option>
            <option value="0">mToken wei</option>
          </select>
          <button>Propose</button>
        </form>

		<h3>Mint operations</h3>

        <MintList entries={this.props.mintList || []} governor={this.state.governor} sendData={this.props.sendData}/>
      </fieldset>
    )
  }
  handleGovernorChange(e) {
    this.setState({governor: e.target.value})
  }
  handleAddressChange(e) {
    this.setState({address: e.target.value})
  }
  handleAmountChange(e) {
    this.setState({amount: e.target.value})
  }
  handleUnitChange(e) {
    this.setState({unit: e.target.value})
  }

  handleProposeMint(e) {
    e.preventDefault();
    this.props.sendData({
      action: "mint",
      governor: this.state.governor,
      mint_address: this.state.address,
      mint_amount: this.state.amount,
      mint_unit: this.state.unit
    })
  }
}

class App extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      accounts: [],
      coinbase: "",
      block: 0,
    };
  }

  componentDidMount() {
    this.reconnect()
  }
  
  componentWillUnmount() {
    this.ws.close();
    this.ws = undefined;
  }

  reconnect() {
    this.ws = new WebSocket(((window.location.protocol === "https:") ? "wss://" : "ws://") + window.location.host + "/api");
    this.ws.onclose = () => {
      setTimeout(this.reconnect.bind(this), 3000);
    };
    this.ws.onmessage = this.onWsMessage.bind(this)
  }

  onWsMessage(event) {
    var msg = JSON.parse(event.data);
    if (msg === null) {
      return;
    }

    if (msg.block !== undefined) {
      this.setState({block: msg.block});
    }
    if (msg.coinbase !== undefined) {
      this.setState({coinbase: msg.coinbase});
    }
    if (msg.accounts !== undefined) {
      this.setState({accounts: msg.accounts});
    }
    if (msg.mintList !== undefined) {
      this.setState({mintList: msg.mintList});
    }
    if (msg.error !== undefined) {
      noty({layout: 'topCenter', text: msg.error, type: 'error', timeout: 5000, progressBar: true});
    }
    if (msg.success !== undefined) {
      noty({layout: 'topCenter', text: msg.success, type: 'success', timeout: 5000, progressBar: true});
    }
  }

  sendData(data) {
    this.ws.send(JSON.stringify(data));
  }

  render() {
    return (
      <div class="container">
	<header>
        <div class="row">
          <div class="col-lg-12">
            <h1>kcoin Control Panel</h1>
          </div>
        </div>
        <div class="row summary">
          <div class="col-sm-5"><i class="fa fa-database" aria-hidden="true"></i> {this.state.coinbase} coinbase</div>
          <div class="col-sm-3"><i class="fa fa-database" aria-hidden="true"></i> {this.state.block} blocks</div>
        </div>
	</header>

        <Minting accounts={this.state.accounts} mintList={this.state.mintList} sendData={this.sendData.bind(this)} />
      </div>
    )
  }
}

ReactDOM.render(
  <App />,
  document.getElementById('control')
);


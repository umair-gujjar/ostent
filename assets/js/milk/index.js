// Generated by CoffeeScript 1.7.1
(function() {
  this.newwebsocket = function(onmessage) {
    var conn, init, sendClient, sendJSON, sendSearch;
    conn = null;
    sendSearch = function(search) {
      return sendJSON({
        Search: search
      });
    };
    sendClient = function(client) {
      console.log(JSON.stringify(client), 'sendClient');
      return sendJSON({
        Client: client
      });
    };
    sendJSON = function(obj) {
      if ((conn == null) || conn.readyState === conn.CLOSING || conn.readyState === conn.CLOSED) {
        init();
      }
      if ((conn == null) || conn.readyState !== conn.OPEN) {
        console.log('Not connected, cannot send', obj);
        return;
      }
      return conn.send(JSON.stringify(obj));
    };
    init = function() {
      var again, hostport, statesel;
      hostport = window.location.hostname + (location.port ? ':' + location.port : '');
      conn = new WebSocket('ws://' + hostport + '/ws');
      conn.onopen = function() {
        sendSearch(location.search);
        $(window).bind('popstate', (function() {
          sendSearch(location.search);
        }));
      };
      statesel = 'table thead tr .header a.state';
      again = function(e) {
        $(statesel).unbind('click');
        if (!e.wasClean) {
          window.setTimeout(init, 5000);
        }
      };
      conn.onclose = again;
      conn.onerror = again;
      conn.onmessage = onmessage;
      $(statesel).click(function() {
        history.pushState({
          path: this.path
        }, '', this.href);
        sendSearch(this.search);
        return false;
      });
    };
    init();
    return {
      sendClient: sendClient,
      sendSearch: sendSearch,
      close: function() {
        return conn.close();
      }
    };
  };

  this.IFbytesCLASS = React.createClass({
    getInitialState: function() {
      return Data.IFbytes;
    },
    render: function() {
      var $if, Data;
      Data = {
        IFbytes: this.state
      };
      return ifbytes_table(Data, (function() {
        var _i, _len, _ref, _ref1, _ref2, _results;
        _ref2 = (_ref = Data != null ? (_ref1 = Data.IFbytes) != null ? _ref1.List : void 0 : void 0) != null ? _ref : [];
        _results = [];
        for (_i = 0, _len = _ref2.length; _i < _len; _i++) {
          $if = _ref2[_i];
          _results.push(ifbytes_rows(Data, $if));
        }
        return _results;
      })());
    }
  });

  this.IFerrorsCLASS = React.createClass({
    getInitialState: function() {
      return Data.IFerrors;
    },
    render: function() {
      var $if, Data;
      Data = {
        IFerrors: this.state
      };
      return iferrors_table(Data, (function() {
        var _i, _len, _ref, _ref1, _ref2, _results;
        _ref2 = (_ref = Data != null ? (_ref1 = Data.IFerrors) != null ? _ref1.List : void 0 : void 0) != null ? _ref : [];
        _results = [];
        for (_i = 0, _len = _ref2.length; _i < _len; _i++) {
          $if = _ref2[_i];
          _results.push(iferrors_rows(Data, $if));
        }
        return _results;
      })());
    }
  });

  this.IFpacketsCLASS = React.createClass({
    getInitialState: function() {
      return Data.IFpackets;
    },
    render: function() {
      var $if, Data;
      Data = {
        IFpackets: this.state
      };
      return ifpackets_table(Data, (function() {
        var _i, _len, _ref, _ref1, _ref2, _results;
        _ref2 = (_ref = Data != null ? (_ref1 = Data.IFpackets) != null ? _ref1.List : void 0 : void 0) != null ? _ref : [];
        _results = [];
        for (_i = 0, _len = _ref2.length; _i < _len; _i++) {
          $if = _ref2[_i];
          _results.push(ifpackets_rows(Data, $if));
        }
        return _results;
      })());
    }
  });

  this.DFbytesCLASS = React.createClass({
    getInitialState: function() {
      return {
        DFlinks: Data.DFlinks,
        DFbytes: Data.DFbytes
      };
    },
    render: function() {
      var $disk, Data;
      Data = this.state;
      return dfbytes_table(Data, (function() {
        var _i, _len, _ref, _ref1, _ref2, _results;
        _ref2 = (_ref = Data != null ? (_ref1 = Data.DFbytes) != null ? _ref1.List : void 0 : void 0) != null ? _ref : [];
        _results = [];
        for (_i = 0, _len = _ref2.length; _i < _len; _i++) {
          $disk = _ref2[_i];
          _results.push(dfbytes_rows(Data, $disk));
        }
        return _results;
      })());
    }
  });

  this.DFinodesCLASS = React.createClass({
    getInitialState: function() {
      return {
        DFlinks: Data.DFlinks,
        DFinodes: Data.DFinodes
      };
    },
    render: function() {
      var $disk, Data;
      Data = this.state;
      return dfinodes_table(Data, (function() {
        var _i, _len, _ref, _ref1, _ref2, _results;
        _ref2 = (_ref = Data != null ? (_ref1 = Data.DFinodes) != null ? _ref1.List : void 0 : void 0) != null ? _ref : [];
        _results = [];
        for (_i = 0, _len = _ref2.length; _i < _len; _i++) {
          $disk = _ref2[_i];
          _results.push(dfinodes_rows(Data, $disk));
        }
        return _results;
      })());
    }
  });

  this.MEMtableCLASS = React.createClass({
    getInitialState: function() {
      return Data.MEM;
    },
    render: function() {
      var $mem, Data;
      Data = {
        MEM: this.state
      };
      return mem_table(Data, (function() {
        var _i, _len, _ref, _ref1, _ref2, _results;
        _ref2 = (_ref = Data != null ? (_ref1 = Data.MEM) != null ? _ref1.List : void 0 : void 0) != null ? _ref : [];
        _results = [];
        for (_i = 0, _len = _ref2.length; _i < _len; _i++) {
          $mem = _ref2[_i];
          _results.push(mem_rows(Data, $mem));
        }
        return _results;
      })());
    }
  });

  this.CPUtableCLASS = React.createClass({
    getInitialState: function() {
      return Data.CPU;
    },
    render: function() {
      var $core, Data;
      Data = {
        CPU: this.state
      };
      return cpu_table(Data, (function() {
        var _i, _len, _ref, _ref1, _ref2, _results;
        _ref2 = (_ref = Data != null ? (_ref1 = Data.CPU) != null ? _ref1.List : void 0 : void 0) != null ? _ref : [];
        _results = [];
        for (_i = 0, _len = _ref2.length; _i < _len; _i++) {
          $core = _ref2[_i];
          _results.push(cpu_rows(Data, $core));
        }
        return _results;
      })());
    }
  });

  this.PStableCLASS = React.createClass({
    getInitialState: function() {
      return {
        PStable: Data.PStable,
        PSlinks: Data.PSlinks
      };
    },
    render: function() {
      var $proc, Data;
      Data = this.state;
      return ps_table(Data, (function() {
        var _i, _len, _ref, _ref1, _ref2, _results;
        _ref2 = (_ref = Data != null ? (_ref1 = Data.PStable) != null ? _ref1.List : void 0 : void 0) != null ? _ref : [];
        _results = [];
        for (_i = 0, _len = _ref2.length; _i < _len; _i++) {
          $proc = _ref2[_i];
          _results.push(ps_rows(Data, $proc));
        }
        return _results;
      })());
    }
  });

  this.VGtableCLASS = React.createClass({
    getInitialState: function() {
      return {
        VagrantMachines: Data.VagrantMachines,
        VagrantError: Data.VagrantError,
        VagrantErrord: Data.VagrantErrord
      };
    },
    render: function() {
      var $machine, Data, rows;
      Data = this.state;
      if (((Data != null ? Data.VagrantErrord : void 0) != null) && Data.VagrantErrord) {
        rows = [vagrant_error(Data)];
      } else {
        rows = (function() {
          var _i, _len, _ref, _ref1, _ref2, _results;
          _ref2 = (_ref = Data != null ? (_ref1 = Data.VagrantMachines) != null ? _ref1.List : void 0 : void 0) != null ? _ref : [];
          _results = [];
          for (_i = 0, _len = _ref2.length; _i < _len; _i++) {
            $machine = _ref2[_i];
            _results.push(vagrant_rows(Data, $machine));
          }
          return _results;
        })();
      }
      return vagrant_table(Data, rows);
    }
  });

  this.HideClass = React.createClass({
    reduce: function(data) {
      var value;
      if ((data != null ? data.Client : void 0) != null) {
        value = data.Client[this.props.key];
        if (value !== void 0) {
          return {
            Hide: value
          };
        }
      }
    },
    getInitialState: function() {
      return this.reduce(Data);
    },
    componentDidMount: function() {
      return this.props.$click_el.click(this.click);
    },
    render: function() {
      this.props.$collapse_el.collapse(this.state.Hide ? 'hide' : 'show');
      return React.DOM.span();
    },
    click: function(e) {
      var S;
      (S = {})[this.props.key] = !this.state.Hide;
      websocket.sendClient(S);
      e.stopPropagation();
      e.preventDefault();
      return void 0;
    }
  });

  this.ShowSwapClass = React.createClass({
    getInitialState: function() {
      return ShowSwapClass.reduce(Data);
    },
    statics: {
      reduce: function(data) {
        var S;
        if ((data != null ? data.Client : void 0) != null) {
          S = {};
          if (data.Client.HideSWAP !== void 0) {
            S.HideSWAP = data.Client.HideSWAP;
          }
          if (data.Client.HideMEM !== void 0) {
            S.HideMEM = data.Client.HideMEM;
          }
          return S;
        }
      }
    },
    componentDidMount: function() {
      return this.props.$el.click(this.click);
    },
    render: function() {
      this.props.$el[!this.state.HideSWAP ? 'addClass' : 'removeClass']('active');
      return React.DOM.span(null, this.props.$el.text());
    },
    click: function(e) {
      var S;
      S = {
        HideSWAP: !this.state.HideSWAP
      };
      if (this.state.HideMEM) {
        S.HideMEM = false;
      }
      websocket.sendClient(S);
      e.stopPropagation();
      e.preventDefault();
      return void 0;
    }
  });

  this.NewTextCLASS = function(reduce) {
    return React.createClass({
      newstate: function(data) {
        var v;
        v = reduce(data);
        if (v != null) {
          return {
            Text: v
          };
        }
      },
      getInitialState: function() {
        return this.newstate(Data);
      },
      render: function() {
        return React.DOM.span(null, this.state.Text);
      }
    });
  };

  this.setState = function(obj, data) {
    var key;
    if (data != null) {
      for (key in data) {
        if (data[key] == null) {
          delete data[key];
        }
      }
      return obj.setState(data);
    }
  };

  this.dummy = function(sel) {
    return sel.append('<span class="dummy display-none" />').find('.dummy').get(0);
  };

  this.update = function(currentClient, model) {
    var $header_mem, $hiding_mem, $showswap_el, cputable, dfbytes, dfinodes, dftitle, hideconfigmem, hidemem, hostname, ifbytes, iferrors, ifpackets, iftitle, ip, la, memtable, onmessage, param, psplus, pstable, showswap, uptime, vagrant;
    if (((function() {
      var _i, _len, _ref, _results;
      _ref = location.search.substr(1).split('&');
      _results = [];
      for (_i = 0, _len = _ref.length; _i < _len; _i++) {
        param = _ref[_i];
        if (param.split('=')[0] === 'still') {
          _results.push(42);
        }
      }
      return _results;
    })()).length) {
      return;
    }
    $showswap_el = $('label[href="#showswap"]');
    showswap = React.renderComponent(ShowSwapClass({
      $el: $showswap_el
    }), $showswap_el.get(0));
    $header_mem = $('header a[href="#mem"]');
    hideconfigmem = React.renderComponent(HideClass({
      key: 'HideconfigMEM',
      $collapse_el: $('#memconfig'),
      $click_el: $header_mem
    }), dummy($header_mem));
    $hiding_mem = $('#memconfig').find('.hiding');
    hidemem = React.renderComponent(HideClass({
      key: 'HideMEM',
      $collapse_el: $('#mem'),
      $click_el: $hiding_mem
    }), dummy($hiding_mem));
    ip = React.renderComponent(NewTextCLASS(function(data) {
      var _ref;
      return data != null ? (_ref = data.Generic) != null ? _ref.IP : void 0 : void 0;
    })(), $('#generic-ip').get(0));
    hostname = React.renderComponent(NewTextCLASS(function(data) {
      var _ref;
      return data != null ? (_ref = data.Generic) != null ? _ref.Hostname : void 0 : void 0;
    })(), $('#generic-hostname').get(0));
    uptime = React.renderComponent(NewTextCLASS(function(data) {
      var _ref;
      return data != null ? (_ref = data.Generic) != null ? _ref.Uptime : void 0 : void 0;
    })(), $('#generic-uptime').get(0));
    la = React.renderComponent(NewTextCLASS(function(data) {
      var _ref;
      return data != null ? (_ref = data.Generic) != null ? _ref.LA : void 0 : void 0;
    })(), $('#generic-la').get(0));
    iftitle = React.renderComponent(NewTextCLASS(function(data) {
      var _ref;
      return data != null ? (_ref = data.Client) != null ? _ref.TabTitleIF : void 0 : void 0;
    })(), $('header a[href="#if"]').get(0));
    dftitle = React.renderComponent(NewTextCLASS(function(data) {
      var _ref;
      return data != null ? (_ref = data.Client) != null ? _ref.TabTitleDF : void 0 : void 0;
    })(), $('header a[href="#df"]').get(0));
    psplus = React.renderComponent(NewTextCLASS(function(data) {
      var _ref;
      return data != null ? (_ref = data.Client) != null ? _ref.PSplusText : void 0 : void 0;
    })(), $('label.more[href="#psmore"]').get(0));
    memtable = React.renderComponent(MEMtableCLASS(), document.getElementById('mem' + '-' + 'table'));
    pstable = React.renderComponent(PStableCLASS(), document.getElementById('ps' + '-' + 'table'));
    dfbytes = React.renderComponent(DFbytesCLASS(), document.getElementById('dfbytes' + '-' + 'table'));
    dfinodes = React.renderComponent(DFinodesCLASS(), document.getElementById('dfinodes' + '-' + 'table'));
    cputable = React.renderComponent(CPUtableCLASS(), document.getElementById('cpu' + '-' + 'table'));
    ifbytes = React.renderComponent(IFbytesCLASS(), document.getElementById('ifbytes' + '-' + 'table'));
    iferrors = React.renderComponent(IFerrorsCLASS(), document.getElementById('iferrors' + '-' + 'table'));
    ifpackets = React.renderComponent(IFpacketsCLASS(), document.getElementById('ifpackets' + '-' + 'table'));
    vagrant = React.renderComponent(VGtableCLASS(), document.getElementById('vagrant' + '-' + 'table'));
    onmessage = function(event) {
      var data, _ref;
      data = JSON.parse(event.data);
      if (data == null) {
        return;
      }
      if (((_ref = data.Client) != null ? _ref.DebugError : void 0) != null) {
        console.log('DEBUG ERROR', data.Client.DebugError);
      }
      if ((data.Reload != null) && data.Reload) {
        window.setTimeout((function() {
          return location.reload(true);
        }), 5000);
        window.setTimeout(websocket.close, 2000);
        console.log('in 5s: location.reload(true)');
        console.log('in 2s: websocket.close()');
        return;
      }
      setState(pstable, {
        PStable: data.PStable,
        PSlinks: data.PSlinks
      });
      setState(dfbytes, {
        DFbytes: data.DFbytes,
        DFlinks: data.DFlinks
      });
      setState(dfinodes, {
        DFinodes: data.DFinodes,
        DFlinks: data.DFlinks
      });
      setState(showswap, ShowSwapClass.reduce(data));
      setState(hideconfigmem, hideconfigmem.reduce(data));
      setState(hidemem, hidemem.reduce(data));
      setState(ip, ip.newstate(data));
      setState(hostname, hostname.newstate(data));
      setState(uptime, uptime.newstate(data));
      setState(la, la.newstate(data));
      setState(iftitle, iftitle.newstate(data));
      setState(dftitle, iftitle.newstate(data));
      setState(psplus, psplus.newstate(data));
      setState(memtable, data.MEM);
      setState(cputable, data.CPU);
      setState(ifbytes, data.IFbytes);
      setState(iferrors, data.IFerrors);
      setState(ifpackets, data.IFpackets);
      setState(vagrant, {
        VagrantMachines: data.VagrantMachines,
        VagrantError: data.VagrantError,
        VagrantErrord: data.VagrantErrord
      });
      if (data.Client != null) {
        console.log(JSON.stringify(data.Client), 'recvClient');
      }
      if (data.Client != null) {
        currentClient = React.addons.update(currentClient, {
          $merge: data.Client
        });
      }
      data.Client = currentClient;
      model.set(Model.attributes(data));
      $('span .tooltipable').popover({
        trigger: 'hover focus'
      });
      return $('span .tooltipabledots').popover();
    };
    this.websocket = newwebsocket(onmessage);
  };

  this.Model = Backbone.Model.extend({});

  this.Model.attributes = function(data) {
    if (data.Generic == null) {
      return data.Client;
    }
    if (data.Client == null) {
      return data.Generic;
    }
    return React.addons.update(data.Generic, {
      $merge: data.Client
    });
  };

  this.View = Backbone.View.extend({
    initialize: function() {
      var $config_cpu, $config_df, $config_if, $config_mem, $config_ps, $config_vg, $header_cpu, $header_df, $header_if, $header_ps, $header_vg, $hidden_cpu, $hidden_df, $hidden_if, $hidden_ps, $hidden_vg, $panels_df, $panels_if, $psless, $psmore, $section_cpu, $section_df, $section_if, $section_ps, $section_vg, $tab_df, $tab_if, B, doexpandable, expandable_sections, sections, _i, _len;
      $section_if = $('#if');
      $section_cpu = $('#cpu');
      $section_df = $('#df');
      $section_ps = $('#ps');
      $section_vg = $('#vagrant');
      $config_if = $('#ifconfig');
      $config_cpu = $('#cpuconfig');
      $config_df = $('#dfconfig');
      $config_ps = $('#psconfig');
      $config_vg = $('#vgconfig');
      $hidden_if = $config_if.find('.hiding');
      $hidden_cpu = $config_cpu.find('.hiding');
      $hidden_df = $config_df.find('.hiding');
      $hidden_ps = $config_ps.find('.hiding');
      $hidden_vg = $config_vg.find('.hiding');
      this.listenhide('HideCPU', $section_cpu, $hidden_cpu);
      this.listenhide('HidePS', $section_ps, $hidden_ps);
      this.listenhide('HideVG', $section_vg, $hidden_vg);
      $header_if = $('header a[href="' + $section_if.selector + '"]');
      $header_cpu = $('header a[href="' + $section_cpu.selector + '"]');
      $header_df = $('header a[href="' + $section_df.selector + '"]');
      $header_ps = $('header a[href="' + $section_ps.selector + '"]');
      $header_vg = $('header a[href="' + $section_vg.selector + '"]');
      this.listenhide('HideconfigIF', $config_if, $header_if, true);
      this.listenhide('HideconfigCPU', $config_cpu, $header_cpu, true);
      this.listenhide('HideconfigDF', $config_df, $header_df, true);
      this.listenhide('HideconfigPS', $config_ps, $header_ps, true);
      this.listenhide('HideconfigVG', $config_vg, $header_vg, true);
      $tab_if = $('.if-switch');
      $tab_df = $('.df-switch');
      $panels_if = $('.if-tab');
      $panels_df = $('.df-tab');
      this.listenTo(this.model, 'change:HideIF', this.change_collapsetabfunc('HideIF', 'TabIF', $panels_if, $tab_if));
      this.listenTo(this.model, 'change:HideDF', this.change_collapsetabfunc('HideDF', 'TabDF', $panels_df, $tab_df));
      this.listenTo(this.model, 'change:TabIF', this.change_collapsetabfunc('HideIF', 'TabIF', $panels_if, $tab_if));
      this.listenTo(this.model, 'change:TabDF', this.change_collapsetabfunc('HideDF', 'TabDF', $panels_df, $tab_df));
      $psmore = $('label.more[href="#psmore"]');
      $psless = $('label.less[href="#psless"]');
      this.listenenable('PSnotExpandable', $psmore);
      this.listenenable('PSnotDecreasable', $psless);
      $config_mem = $('#memconfig');
      this.listenrefresherror('RefreshErrorMEM', $config_mem.find('.refresh-group'));
      this.listenrefresherror('RefreshErrorIF', $config_if.find('.refresh-group'));
      this.listenrefresherror('RefreshErrorCPU', $config_cpu.find('.refresh-group'));
      this.listenrefresherror('RefreshErrorDF', $config_df.find('.refresh-group'));
      this.listenrefresherror('RefreshErrorPS', $config_ps.find('.refresh-group'));
      this.listenrefresherror('RefreshErrorVG', $config_vg.find('.refresh-group'));
      this.listenrefreshvalue('RefreshMEM', $config_mem.find('.refresh-input'));
      this.listenrefreshvalue('RefreshIF', $config_if.find('.refresh-input'));
      this.listenrefreshvalue('RefreshCPU', $config_cpu.find('.refresh-input'));
      this.listenrefreshvalue('RefreshDF', $config_df.find('.refresh-input'));
      this.listenrefreshvalue('RefreshPS', $config_ps.find('.refresh-input'));
      this.listenrefreshvalue('RefreshVG', $config_vg.find('.refresh-input'));
      B = function(c) {
        return c;
      };
      expandable_sections = [[$section_if, 'ExpandIF', 'HideIF', 'ExpandableIF', 'ExpandtextIF'], [$section_cpu, 'ExpandCPU', 'HideCPU', 'ExpandableCPU', 'ExpandtextCPU'], [$section_df, 'ExpandDF', 'HideDF', 'ExpandableDF', 'ExpandtextDF']];
      doexpandable = (function(_this) {
        return function(sections) {
          var $b, E, H, L, S, T;
          S = sections[0];
          E = sections[1];
          H = sections[2];
          L = sections[3];
          T = sections[4];
          $b = $('label[href="' + S.selector + '"]');
          _this.listentext(T, $b);
          _this.listenenable(L, $b, true);
          _this.listenactivate(E, $b);
          $b.click(B(_this.click_expandfunc(E, H)));
        };
      })(this);
      for (_i = 0, _len = expandable_sections.length; _i < _len; _i++) {
        sections = expandable_sections[_i];
        doexpandable(sections);
      }
      $tab_if.click(B(this.click_tabfunc('TabIF', 'HideIF')));
      $tab_df.click(B(this.click_tabfunc('TabDF', 'HideDF')));
      $header_if.click(B(this.click_expandfunc('HideconfigIF')));
      $header_cpu.click(B(this.click_expandfunc('HideconfigCPU')));
      $header_df.click(B(this.click_expandfunc('HideconfigDF')));
      $header_ps.click(B(this.click_expandfunc('HideconfigPS')));
      $header_vg.click(B(this.click_expandfunc('HideconfigVG')));
      $hidden_if.click(B(this.click_expandfunc('HideIF')));
      $hidden_cpu.click(B(this.click_expandfunc('HideCPU')));
      $hidden_df.click(B(this.click_expandfunc('HideDF')));
      $hidden_ps.click(B(this.click_expandfunc('HidePS')));
      $hidden_vg.click(B(this.click_expandfunc('HideVG')));
      $psmore.click(B(this.click_psignalfunc('HidePS', true)));
      $psless.click(B(this.click_psignalfunc('HidePS', false)));
      $config_mem.find('.refresh-input').on('input', B(this.submit_rsignalfunc('RefreshSignalMEM')));
      $config_if.find('.refresh-input').on('input', B(this.submit_rsignalfunc('RefreshSignalIF')));
      $config_cpu.find('.refresh-input').on('input', B(this.submit_rsignalfunc('RefreshSignalCPU')));
      $config_df.find('.refresh-input').on('input', B(this.submit_rsignalfunc('RefreshSignalDF')));
      $config_ps.find('.refresh-input').on('input', B(this.submit_rsignalfunc('RefreshSignalPS')));
      $config_vg.find('.refresh-input').on('input', B(this.submit_rsignalfunc('RefreshSignalVG')));
    },
    submit_rsignalfunc: function(R) {
      return function(e) {
        var S;
        (S = {})[R] = $(e.target).val();
        websocket.sendClient(S);
      };
    },
    listentext: function(K, $el) {
      return this.listenTo(this.model, 'change:' + K, this._text(K, $el));
    },
    _text: function(K, $el) {
      return function() {
        return $el.text(this.model.attributes[K]);
      };
    },
    listenrefresherror: function(E, $el) {
      return this.listenTo(this.model, 'change:' + E, function() {
        return $el[this.model.attributes[E] ? 'addClass' : 'removeClass']('has-warning');
      });
    },
    listenrefreshvalue: function(E, $el) {
      return this.listenTo(this.model, 'change:' + E, function() {
        return $el.prop('value', this.model.attributes[E]);
      });
    },
    listenenable: function(K, $el, reverse) {
      return this.listenTo(this.model, 'change:' + K, function() {
        var V;
        V = (this.model.attributes[K] != null) && this.model.attributes[K];
        if ((reverse != null) && reverse) {
          V = !V;
        }
        $el.prop('disabled', V);
        return $el[V ? 'addClass' : 'removeClass']('disabled');
      });
    },
    listenactivate: function(K, $el, reverse) {
      return this.listenTo(this.model, 'change:' + K, function() {
        var V;
        V = (this.model.attributes[K] != null) && this.model.attributes[K];
        if ((reverse == null) && reverse) {
          V = !V;
        }
        return $el[V ? 'addClass' : 'removeClass']('active');
      });
    },
    listenhide: function(H, $el, $button_el, reverse) {
      return this.listenTo(this.model, 'change:' + H, function() {
        var V;
        V = (this.model.attributes[H] != null) && this.model.attributes[H];
        $el.collapse(V ? 'hide' : 'show');
        if ((reverse == null) && reverse) {
          V = !V;
        }
        return $button_el[V ? 'addClass' : 'removeClass']('active');
      });
    },
    change_collapsefunc: function(H, $el) {
      return function() {
        return $el.collapse(this.model.attributes[H] ? 'hide' : 'show');
      };
    },
    change_collapsetabfunc: function(H, T, $el, $tabel) {
      return function() {
        var A, activeClass, curtabid, el, nots, _i, _j, _len, _len1;
        A = this.model.attributes;
        if (A[H]) {
          $el.collapse('hide');
          return;
        }
        curtabid = A[T];
        nots = $el.not('[data-tabid="' + curtabid + '"]');
        for (_i = 0, _len = nots.length; _i < _len; _i++) {
          el = nots[_i];
          $(el).collapse('hide');
        }
        $($el.not(nots)).collapse('show');
        activeClass = function(el) {
          var tabid_attr, xel;
          xel = $(el);
          tabid_attr = +xel.attr('data-tabid');
          return xel[tabid_attr === curtabid ? 'addClass' : 'removeClass']('active');
        };
        for (_j = 0, _len1 = $tabel.length; _j < _len1; _j++) {
          el = $tabel[_j];
          activeClass(el);
        }
      };
    },
    click_expandfunc: function(H, H2) {
      return (function(_this) {
        return function(e) {
          var A, S;
          A = _this.model.attributes;
          (S = {})[H] = !A[H];
          if ((H2 != null) && A[H2]) {
            S[H2] = !A[H2];
          }
          websocket.sendClient(S);
          e.preventDefault();
          return e.stopPropagation();
        };
      })(this);
    },
    click_tabfunc: function(T, H) {
      return (function(_this) {
        return function(e) {
          var S, V, newtabid;
          newtabid = +$($(e.target).attr('href')).attr('data-tabid');
          (S = {})[T] = newtabid;
          V = _this.model.attributes[H];
          if (V) {
            S[H] = !V;
          }
          websocket.sendClient(S);
          e.preventDefault();
          return e.stopPropagation();
        };
      })(this);
    },
    click_psignalfunc: function(H, v) {
      return (function(_this) {
        return function(e) {
          var S, V;
          S = {
            MorePsignal: v
          };
          V = _this.model.attributes[H];
          if (V) {
            S[H] = !V;
          }
          websocket.sendClient(S);
          e.preventDefault();
          return e.stopPropagation();
        };
      })(this);
    }
  });

  this.ready = function() {
    var model;
    (new Headroom(document.querySelector('nav'), {
      offset: 71 - 51
    })).init();
    $('.collapse').collapse({
      toggle: false
    });
    $('span .tooltipable').popover({
      trigger: 'hover focus'
    });
    $('span .tooltipabledots').popover();
    $('[data-toggle="popover"]').popover();
    $('#generic-la').popover({
      trigger: 'hover focus',
      placement: 'right',
      html: true,
      content: function() {
        return $('#uptime').html();
      }
    });
    $('body').on('click', function(e) {
      $('span .tooltipabledots').each(function() {
        if (!$(this).is(e.target) && $(this).has(e.target).length === 0 && $('.popover').has(e.target).length === 0) {
          $(this).popover('hide');
        }
      });
    });
    model = new Model(Model.attributes(Data));
    new View({
      model: model
    });
    update(Data.Client, model);
  };

}).call(this);

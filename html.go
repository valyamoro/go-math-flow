package main

const htmlTemplate = `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width,initial-scale=1">
  <title>Linear Equation</title>
  <link rel="preconnect" href="https://fonts.googleapis.com">
  <link href="https://fonts.googleapis.com/css2?family=Space+Mono:wght@400;700&family=Oxanium:wght@300;400;600;800&display=swap" rel="stylesheet">
  <script src="https://cdn.plot.ly/plotly-2.32.0.min.js" charset="utf-8"></script>
  <style>
    :root {
      --bg:        #06101e;
      --surface:   #0a1628;
      --surface2:  #0e1f38;
      --border:    rgba(100,160,255,0.1);
      --border2:   rgba(100,160,255,0.2);
      --text:      #c5d8f0;
      --muted:     #4a6a9a;
      --muted2:    #6a8ab0;
      --accent:    #00e6a0;
      --accent2:   #00b87a;
      --blue:      #4d9eff;
      --red:       #ff4d6d;
      --yellow:    #ffd166;
    }

    *, *::before, *::after { box-sizing: border-box; margin: 0; padding: 0; }

    body {
      background: var(--bg);
      color: var(--text);
      font-family: 'Space Mono', 'Courier New', monospace;
      min-height: 100vh;
      display: flex;
      flex-direction: column;
      overflow: hidden;
    }

    body::before {
      content: '';
      position: fixed; inset: 0;
      background: repeating-linear-gradient(
        0deg,
        transparent,
        transparent 2px,
        rgba(0,0,0,0.04) 2px,
        rgba(0,0,0,0.04) 4px
      );
      pointer-events: none;
      z-index: 1000;
    }

    header {
      background: var(--surface);
      border-bottom: 1px solid var(--border2);
      padding: 0 24px;
      height: 56px;
      display: flex;
      align-items: center;
      gap: 16px;
      flex-shrink: 0;
      position: relative;
      overflow: hidden;
    }
    header::after {
      content: '';
      position: absolute;
      bottom: 0; left: 0; right: 0;
      height: 1px;
      background: linear-gradient(90deg, transparent, var(--accent), transparent);
      opacity: 0.6;
    }

    .logo-mark {
      width: 34px; height: 34px;
      border: 1.5px solid var(--accent);
      border-radius: 6px;
      display: flex; align-items: center; justify-content: center;
      font-family: 'Oxanium', sans-serif;
      font-weight: 800;
      font-size: 16px;
      color: var(--accent);
      flex-shrink: 0;
      box-shadow: 0 0 12px rgba(0,230,160,0.2);
      animation: pulse-border 3s ease infinite;
    }
    @keyframes pulse-border {
      0%, 100% { box-shadow: 0 0 8px rgba(0,230,160,0.2); }
      50%       { box-shadow: 0 0 20px rgba(0,230,160,0.4); }
    }

    .header-title {
      font-family: 'Oxanium', sans-serif;
      font-weight: 600;
      font-size: 15px;
      letter-spacing: .05em;
      color: var(--text);
    }
    .header-title span { color: var(--accent); }

    .header-tag {
      margin-left: auto;
      font-size: 10px;
      color: var(--muted);
      letter-spacing: .08em;
      border: 1px solid var(--border);
      padding: 3px 8px;
      border-radius: 3px;
      background: rgba(0,0,0,0.2);
    }

    .workspace {
      flex: 1;
      display: flex;
      min-height: 0;
    }

    aside {
      width: 288px;
      flex-shrink: 0;
      background: var(--surface);
      border-right: 1px solid var(--border);
      display: flex;
      flex-direction: column;
      overflow-y: auto;
      overflow-x: hidden;
    }

    .sidebar-section {
      padding: 18px 16px;
      border-bottom: 1px solid var(--border);
    }
    .sidebar-section:last-child { border-bottom: none; }

    .section-label {
      font-size: 9px;
      font-weight: 700;
      letter-spacing: .15em;
      text-transform: uppercase;
      color: var(--muted);
      margin-bottom: 12px;
      display: flex;
      align-items: center;
      gap: 8px;
    }
    .section-label::after {
      content: '';
      flex: 1;
      height: 1px;
      background: var(--border);
    }

    /* Equation rows */
    .eq-list { display: flex; flex-direction: column; gap: 7px; }
    .eq-row {
      display: flex;
      align-items: center;
      gap: 10px;
      background: var(--surface2);
      border: 1px solid var(--border);
      border-radius: 5px;
      padding: 8px 11px;
      font-size: 12px;
      transition: border-color .2s, background .2s;
      animation: slide-in .35s ease backwards;
    }
    .eq-row:hover {
      background: rgba(0,230,160,0.05);
      border-color: rgba(0,230,160,0.25);
    }
    .eq-row:nth-child(1) { animation-delay: 0.05s; }
    .eq-row:nth-child(2) { animation-delay: 0.10s; }
    .eq-tag {
      font-size: 9px;
      color: var(--muted);
      letter-spacing: .06em;
      min-width: 52px;
    }
    .eq-label { color: var(--text); font-size: 12px; }

    /* Status box */
    .status-box {
      display: flex;
      align-items: center;
      gap: 10px;
      padding: 11px 14px;
      border-radius: 6px;
      font-size: 12px;
      font-family: 'Oxanium', sans-serif;
      font-weight: 600;
      letter-spacing: .03em;
    }
    .status-box.ok {
      background: rgba(0,230,160,0.08);
      border: 1px solid rgba(0,230,160,0.3);
      color: var(--accent);
    }
    .status-box.fail {
      background: rgba(255,77,109,0.08);
      border: 1px solid rgba(255,77,109,0.3);
      color: var(--red);
    }
    .status-box.infinite {
      background: rgba(77,158,255,0.08);
      border: 1px solid rgba(77,158,255,0.3);
      color: var(--blue);
    }
    .status-icon { font-size: 16px; }

    /* Root value */
    .root-box {
      background: var(--surface2);
      border: 1px solid rgba(0,230,160,0.3);
      border-radius: 6px;
      padding: 12px 14px;
      font-size: 15px;
      color: var(--accent);
      font-weight: 700;
      letter-spacing: .04em;
      text-align: center;
    }

    /* Method steps */
    .step-list { display: flex; flex-direction: column; gap: 6px; }
    .step-row {
      display: flex;
      align-items: center;
      gap: 10px;
      background: var(--surface2);
      border: 1px solid var(--border);
      border-radius: 4px;
      padding: 7px 10px;
      animation: slide-in .35s ease backwards;
    }
    .step-row:nth-child(1) { animation-delay: 0.08s; }
    .step-row:nth-child(2) { animation-delay: 0.14s; }
    .step-row:nth-child(3) { animation-delay: 0.20s; }
    .step-num {
      font-size: 10px;
      color: var(--muted);
      min-width: 16px;
      text-align: center;
      flex-shrink: 0;
    }
    .step-text {
      font-size: 11px;
      color: var(--text);
      font-family: 'Space Mono', monospace;
    }

    /* Range */
    .range-row { display: flex; flex-direction: column; gap: 5px; }
    .range-item {
      background: var(--surface2);
      border: 1px solid var(--border);
      border-radius: 4px;
      padding: 6px 10px;
      font-size: 11px;
      color: var(--muted2);
    }
    .range-item strong { color: var(--blue); }

    /* Chart */
    .chart-wrap {
      flex: 1;
      position: relative;
      min-width: 0;
      min-height: 0;
    }
    #plot { width: 100%; height: 100%; }

    @keyframes slide-in {
      from { opacity: 0; transform: translateX(-8px); }
      to   { opacity: 1; transform: translateX(0); }
    }

    ::-webkit-scrollbar { width: 5px; }
    ::-webkit-scrollbar-track { background: transparent; }
    ::-webkit-scrollbar-thumb {
      background: rgba(100,160,255,0.2);
      border-radius: 3px;
    }
    ::-webkit-scrollbar-thumb:hover { background: rgba(100,160,255,0.35); }
  </style>
</head>
<body>

<!-- HEADER -->
<header>
  <div class="logo-mark">=</div>
  <div class="header-title">
    Linear <span>Equation</span>
  </div>
  <div class="header-tag">Plotly.js · 2D · Go</div>
</header>

<!-- WORKSPACE -->
<div class="workspace">

  <!-- SIDEBAR -->
  <aside>

    <!-- Equation -->
    <div class="sidebar-section">
      <div class="section-label">Equation</div>
      <div class="eq-list">
        <div class="eq-row">
          <span class="eq-tag">input</span>
          <span class="eq-label">{{.Eq.Original}}</span>
        </div>
        <div class="eq-row">
          <span class="eq-tag">standard</span>
          <span class="eq-label">{{.Eq.Standard}}</span>
        </div>
      </div>
    </div>

    <!-- Result -->
    <div class="sidebar-section">
      <div class="section-label">Result</div>
      {{if isUnique .Eq.Kind}}
      <div class="status-box ok">
        <span class="status-icon">✓</span>
        One root
      </div>
      {{else if isNone .Eq.Kind}}
      <div class="status-box fail">
        <span class="status-icon">✗</span>
        No solution
      </div>
      {{else}}
      <div class="status-box infinite">
        <span class="status-icon">∞</span>
        All real numbers
      </div>
      {{end}}
    </div>

    <!-- Root value -->
    {{if .Eq.HasRoot}}
    <div class="sidebar-section">
      <div class="section-label">Root</div>
      <div class="root-box">x = {{fmtF .Eq.Root}}</div>
    </div>

    <!-- Method -->
    <div class="sidebar-section">
      <div class="section-label">Method</div>
      <div class="step-list">
        <div class="step-row">
          <span class="step-num">①</span>
          <span class="step-text">{{.Eq.Standard}}</span>
        </div>
        <div class="step-row">
          <span class="step-num">②</span>
          <span class="step-text">{{.Eq.StepMul}}</span>
        </div>
        <div class="step-row">
          <span class="step-num">③</span>
          <span class="step-text">{{.Eq.StepDiv}}</span>
        </div>
      </div>
    </div>
    {{end}}

    <!-- View range -->
    <div class="sidebar-section">
      <div class="section-label">View range</div>
      <div class="range-row">
        <div class="range-item"><strong>x</strong> ∈ [{{fmtF .XMin}}, {{fmtF .XMax}}]</div>
        <div class="range-item"><strong>y</strong> ∈ [{{fmtF .YMin}}, {{fmtF .YMax}}]</div>
      </div>
    </div>

  </aside>

  <!-- CHART -->
  <div class="chart-wrap">
    <div id="plot"></div>
  </div>

</div>

<script>
(function () {
  var traces = {{.TracesJSON}};
  var layout = {{.LayoutJSON}};
  var config = {
    responsive: true,
    displaylogo: false,
    modeBarButtonsToRemove: ['lasso2d', 'select2d', 'autoScale2d'],
    toImageButtonOptions: { format: 'png', filename: 'linear-equation', scale: 2 }
  };
  Plotly.newPlot('plot', traces, layout, config);
})();
</script>

</body>
</html>`

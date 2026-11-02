package ui

const Page = `
<html lang="en">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>Smart Pharmacy</title>
<style>
  :root {
    --bg: #E8ECF0;
    --bg-soft: #DDE3EA;
    --card: #EEF1F5;
    --text: #1A2332;
    --muted: #6B7A92;
    --line: rgba(180, 190, 210, 0.45);
    --shadow-out: 6px 6px 14px rgba(150, 165, 190, 0.30), -6px -6px 14px rgba(255, 255, 255, 0.88);
    --shadow-in: inset 4px 4px 10px rgba(150, 165, 190, 0.22), inset -4px -4px 10px rgba(255, 255, 255, 0.92);
    --shadow-card: 8px 8px 20px rgba(140, 158, 185, 0.28), -8px -8px 20px rgba(255, 255, 255, 0.90);
    --primary: #008080;
    --primary-light: #20B2AA;
    --success: #16A34A;
    --warn: #D97706;
    --danger: #DC2626;
    --success-soft: #DCFCE7;
    --warn-soft: #FEF3C7;
    --danger-soft: #FEE2E2;
    --sidebar-w: 252px;
    --header-h: 72px;
  }

  * { box-sizing: border-box; margin: 0; padding: 0; }
  html, body {
    height: 100%;
    font-family: "SF Pro Display","SF Pro Text",-apple-system,BlinkMacSystemFont,"Segoe UI",system-ui,sans-serif;
    color: var(--text);
    background: var(--bg);
  }
  body { overflow: hidden; }

  .app { display: flex; height: 100vh; }

  .sidebar {
    width: var(--sidebar-w);
    flex-shrink: 0;
    background: var(--bg);
    border-right: 1px solid var(--line);
    display: flex;
    flex-direction: column;
    overflow-y: auto;
  }

  .brand {
    display: flex;
    align-items: center;
    gap: 0.8rem;
    padding: 1.2rem 1.2rem 1rem;
    border-bottom: 1px solid var(--line);
  }
  .brand-mark {
    width: 40px; height: 40px;
    border-radius: 14px;
    background: linear-gradient(135deg, var(--primary-light), var(--primary));
    box-shadow: var(--shadow-out);
    display: flex;
    align-items: center;
    justify-content: center;
    color: #fff;
    flex-shrink: 0;
  }
  .brand-mark svg { width: 20px; height: 20px; }
  .brand-title {
    font-size: 1.05rem;
    font-weight: 800;
    letter-spacing: -0.02em;
  }
  .brand-title span { color: var(--primary); }

  .nav-group { padding: 1rem 0.9rem 0; }
  .nav-label {
    font-size: 0.68rem;
    font-weight: 800;
    text-transform: uppercase;
    letter-spacing: 0.08em;
    color: var(--muted);
    padding: 0 0.55rem 0.45rem;
  }
  .nav-item {
    display: flex;
    align-items: center;
    gap: 0.72rem;
    padding: 0.78rem 0.9rem;
    margin-bottom: 0.25rem;
    border-radius: 16px;
    cursor: pointer;
    color: var(--muted);
    border: 1px solid transparent;
    transition: all 0.15s ease;
    user-select: none;
  }
  .nav-item:hover { background: rgba(255,255,255,0.55); color: var(--text); }
  .nav-item.active {
    background: rgba(255,255,255,0.85);
    color: var(--primary);
    font-weight: 800;
    box-shadow: var(--shadow-out);
    border-color: var(--line);
  }
  .nav-item.locked {
    opacity: 0.42;
  }
  .nav-icon {
    width: 18px;
    height: 18px;
    flex: 0 0 18px;
    color: currentColor;
  }
  .nav-icon svg {
    width: 18px;
    height: 18px;
    display: block;
  }

  .sidebar-footer {
    margin-top: auto;
    padding: 1rem 0.9rem 1.2rem;
  }

  .main {
    flex: 1;
    min-width: 0;
    display: flex;
    flex-direction: column;
    overflow: hidden;
  }

  .topbar {
    height: var(--header-h);
    flex-shrink: 0;
    display: flex;
    align-items: center;
    gap: 1rem;
    padding: 0 1.4rem;
    border-bottom: 1px solid var(--line);
  }
  .search {
    flex: 1;
    max-width: 420px;
    display: flex;
    align-items: center;
    gap: 0.65rem;
    border-radius: 16px;
    border: 1px solid var(--line);
    background: rgba(255,255,255,0.68);
    box-shadow: var(--shadow-in);
    padding: 0.7rem 0.95rem;
    color: var(--muted);
    font-size: 0.92rem;
  }
  .search svg { width: 16px; height: 16px; flex: 0 0 16px; }

  .topbar-right {
    margin-left: auto;
    display: flex;
    align-items: center;
    gap: 0.65rem;
    flex-wrap: wrap;
    justify-content: flex-end;
  }
  .badge {
    display: inline-flex;
    align-items: center;
    gap: 0.4rem;
    padding: 0.38rem 0.75rem;
    border-radius: 999px;
    border: 1px solid var(--line);
    background: var(--card);
    box-shadow: var(--shadow-out);
    font-size: 0.78rem;
    font-weight: 800;
    white-space: nowrap;
  }
  .badge.good { color: var(--success); background: var(--success-soft); border-color: #BBF7D0; }
  .badge.warn { color: var(--warn); background: var(--warn-soft); border-color: #FDE68A; }
  .badge.bad { color: var(--danger); background: var(--danger-soft); border-color: #FECACA; }
  .dot {
    width: 7px;
    height: 7px;
    border-radius: 50%;
    background: currentColor;
    display: inline-block;
    flex: 0 0 7px;
  }

  .user-pill {
    display: flex;
    align-items: center;
    gap: 0.55rem;
    border-radius: 999px;
    border: 1px solid var(--line);
    background: rgba(255,255,255,0.78);
    box-shadow: var(--shadow-out);
    padding: 0.36rem 0.8rem 0.36rem 0.42rem;
  }
  .avatar {
    width: 30px;
    height: 30px;
    border-radius: 50%;
    background: linear-gradient(135deg, var(--primary-light), var(--primary));
    color: #fff;
    display: flex;
    align-items: center;
    justify-content: center;
    font-size: 0.8rem;
    font-weight: 900;
  }
  .user-meta { line-height: 1.05; }
  .user-name { font-size: 0.84rem; font-weight: 800; }
  .user-role { font-size: 0.71rem; color: var(--muted); }

  .content {
    flex: 1;
    overflow-y: auto;
    padding: 1.35rem;
  }

  .banner {
    display: none;
    border-radius: 18px;
    border: 1px solid var(--line);
    box-shadow: var(--shadow-out);
    padding: 0.85rem 1rem;
    margin-bottom: 1rem;
    font-size: 0.9rem;
    font-weight: 700;
  }
  .banner.good { display: block; background: var(--success-soft); color: #14532d; border-color: #BBF7D0; }
  .banner.warn { display: block; background: var(--warn-soft); color: #92400e; border-color: #FDE68A; }
  .banner.bad  { display: block; background: var(--danger-soft); color: #7f1d1d; border-color: #FECACA; }
  .banner.info { display: block; background: #EFF6FF; color: #1e3a8a; border-color: #BFDBFE; }

  .page-head {
    display: flex;
    justify-content: space-between;
    align-items: flex-start;
    gap: 1rem;
    margin-bottom: 1.2rem;
    flex-wrap: wrap;
  }
  .title {
    font-size: 1.85rem;
    line-height: 1.05;
    font-weight: 900;
    letter-spacing: -0.04em;
  }
  .sub {
    margin-top: 0.35rem;
    font-size: 0.92rem;
    color: var(--muted);
  }
  .actions {
    display: flex;
    gap: 0.6rem;
    flex-wrap: wrap;
  }

  .btn {
    border: none;
    cursor: pointer;
    border-radius: 14px;
    padding: 0.72rem 1rem;
    font-family: inherit;
    font-weight: 800;
    display: inline-flex;
    align-items: center;
    gap: 0.45rem;
    transition: transform 0.08s ease, box-shadow 0.08s ease, opacity 0.12s ease;
    user-select: none;
  }
  .btn:active { transform: translateY(1px); }
  .btn.primary {
    background: linear-gradient(160deg, var(--primary-light), var(--primary));
    color: #fff;
    box-shadow: 0 8px 20px rgba(0,128,128,0.28);
  }
  .btn.ghost {
    background: rgba(255,255,255,0.78);
    color: var(--text);
    border: 1px solid var(--line);
    box-shadow: var(--shadow-out);
  }
  .btn.success {
    background: linear-gradient(160deg, #34d399, var(--success));
    color: #fff;
    box-shadow: 0 8px 18px rgba(22,163,74,0.18);
  }
  .btn:disabled {
    opacity: 0.45;
    cursor: not-allowed;
  }

  .grid-4 {
    display: grid;
    grid-template-columns: repeat(4, minmax(0, 1fr));
    gap: 0.9rem;
    margin-bottom: 1rem;
  }
  @media (max-width: 1100px) { .grid-4 { grid-template-columns: repeat(2, minmax(0, 1fr)); } }
  @media (max-width: 700px) { .grid-4 { grid-template-columns: 1fr; } }

  .stat {
    background: var(--card);
    border: 1px solid var(--line);
    border-radius: 22px;
    box-shadow: var(--shadow-card);
    padding: 1.1rem 1.15rem;
  }
  .stat-top {
    display: flex;
    align-items: flex-start;
    justify-content: space-between;
    gap: 0.5rem;
    margin-bottom: 0.55rem;
  }
  .stat-icon {
    width: 42px;
    height: 42px;
    border-radius: 14px;
    background: rgba(255,255,255,0.75);
    border: 1px solid var(--line);
    box-shadow: var(--shadow-out);
    display: flex;
    align-items: center;
    justify-content: center;
    color: var(--primary);
  }
  .stat-icon svg { width: 22px; height: 22px; }
  .stat-chip {
    display: inline-flex;
    align-items: center;
    gap: 0.35rem;
    padding: 0.28rem 0.6rem;
    border-radius: 999px;
    font-size: 0.74rem;
    font-weight: 800;
  }
  .stat-chip.up { color: var(--success); background: var(--success-soft); }
  .stat-chip.warn { color: var(--warn); background: var(--warn-soft); }
  .stat-chip.down { color: var(--danger); background: var(--danger-soft); }
  .stat-label {
    color: var(--muted);
    font-size: 0.8rem;
    font-weight: 700;
    text-transform: uppercase;
    letter-spacing: 0.05em;
    margin-bottom: 0.25rem;
  }
  .stat-value {
    font-size: 1.45rem;
    font-weight: 900;
    letter-spacing: -0.03em;
  }

  .card {
    background: var(--card);
    border: 1px solid var(--line);
    border-radius: 22px;
    box-shadow: var(--shadow-card);
    padding: 1.2rem 1.2rem 1.25rem;
  }
  .card-title {
    font-size: 0.98rem;
    font-weight: 900;
    letter-spacing: -0.02em;
    margin-bottom: 0.9rem;
  }
  .two-col {
    display: grid;
    grid-template-columns: repeat(2, minmax(0, 1fr));
    gap: 0.9rem;
  }
  @media (max-width: 980px) { .two-col { grid-template-columns: 1fr; } }

  .form-grid {
    display: grid;
    grid-template-columns: repeat(12, minmax(0, 1fr));
    gap: 0.75rem;
    align-items: end;
  }
  .f1 { grid-column: span 3; }
  .f2 { grid-column: span 6; }
  .f4 { grid-column: span 4; }
  .f12 { grid-column: span 12; }
  @media (max-width: 760px) { .f1, .f2, .f4, .f12 { grid-column: span 12; } }

  label {
    display: block;
    font-size: 0.72rem;
    font-weight: 900;
    text-transform: uppercase;
    letter-spacing: 0.06em;
    color: var(--muted);
    margin-bottom: 0.35rem;
  }
  input, select {
    width: 100%;
    border: 1px solid rgba(160, 175, 200, 0.30);
    background: rgba(255,255,255,0.72);
    box-shadow: var(--shadow-in);
    border-radius: 14px;
    padding: 0.72rem 0.9rem;
    font-family: inherit;
    font-size: 0.92rem;
    color: var(--text);
    outline: none;
  }
  input:focus, select:focus {
    border-color: rgba(0,128,128,0.55);
    box-shadow: 0 0 0 3px rgba(0,128,128,0.10), var(--shadow-in);
  }
  input:disabled, select:disabled, button:disabled { opacity: 0.48; cursor: not-allowed; }

  .tabs {
    display: flex;
    flex-wrap: wrap;
    gap: 0.5rem;
    margin-bottom: 1rem;
  }
  .tab-btn {
    border: 1px solid var(--line);
    background: var(--bg);
    box-shadow: var(--shadow-out);
    color: var(--muted);
    border-radius: 999px;
    padding: 0.55rem 0.95rem;
    font-weight: 800;
    cursor: pointer;
  }
  .tab-btn.active {
    background: rgba(255,255,255,0.88);
    color: var(--primary);
  }
  .tab-btn.locked {
    opacity: 0.38;
  }

  .panel { display: none; }
  .panel.active { display: block; }

  .table-wrap {
    overflow-x: auto;
    border: 1px solid var(--line);
    border-radius: 18px;
    background: rgba(255,255,255,0.55);
  }
  table { width: 100%; border-collapse: collapse; }
  th, td {
    padding: 0.78rem 0.9rem;
    text-align: left;
    border-bottom: 1px solid rgba(180, 190, 210, 0.18);
    vertical-align: top;
    font-size: 0.88rem;
  }
  th {
    background: rgba(220,228,240,0.54);
    font-size: 0.72rem;
    font-weight: 900;
    text-transform: uppercase;
    letter-spacing: 0.05em;
    color: var(--muted);
  }
  tbody tr:hover td { background: rgba(0,128,128,0.03); }
  tbody tr:last-child td { border-bottom: 0; }
  .empty {
    padding: 1rem 0.2rem;
    color: var(--muted);
    font-style: italic;
    font-size: 0.9rem;
  }

  .status-line {
    display: flex;
    flex-wrap: wrap;
    gap: 0.4rem;
    margin-top: 0.7rem;
  }
  .chip {
    display: inline-flex;
    align-items: center;
    gap: 0.35rem;
    border-radius: 999px;
    padding: 0.28rem 0.62rem;
    font-size: 0.76rem;
    font-weight: 800;
    border: 1px solid var(--line);
    background: rgba(255,255,255,0.72);
  }
  .chip.good { background: var(--success-soft); color: var(--success); border-color: #BBF7D0; }
  .chip.warn { background: var(--warn-soft); color: var(--warn); border-color: #FDE68A; }
  .chip.bad { background: var(--danger-soft); color: var(--danger); border-color: #FECACA; }
  .chip.neutral { color: var(--muted); }

  .overlay {
    position: relative;
  }
  .overlay.locked::after {
    content: 'Central node is offline — database features are disabled.';
    position: absolute;
    inset: 0;
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 1rem;
    text-align: center;
    font-size: 0.92rem;
    font-weight: 900;
    color: #7c2d12;
    background: rgba(238,241,245,0.76);
    backdrop-filter: blur(4px);
    border-radius: 22px;
    border: 1px dashed rgba(124,45,18,0.30);
  }

  .settings-grid {
    display: grid;
    grid-template-columns: 2fr 2fr 1fr auto;
    gap: 0.8rem;
    align-items: end;
  }
  @media (max-width: 920px) { .settings-grid { grid-template-columns: 1fr; } }

  .footer-note {
    margin-top: 0.7rem;
    color: var(--muted);
    font-size: 0.8rem;
    line-height: 1.5;
  }

  .icon-btn {
    width: 42px;
    height: 42px;
    padding: 0;
    display: inline-flex;
    align-items: center;
    justify-content: center;
  }

  @media (max-width: 760px) {
    .sidebar { display: none; }
    .content { padding: 1rem; }
  }
</style>
</head>
<body>
<div class="app">
  <aside class="sidebar">
    <div class="brand">
      <div class="brand-mark"><svg fill="currentColor" aria-hidden="true" focusable="false" version="1.1" id="Capa_1" xmlns="http://www.w3.org/2000/svg" 
	 viewBox="0 0 399.947 399.947"
	>
<g>
	<path d="M365.226,34.693C342.854,12.321,313.108,0,281.468,0s-61.386,12.321-83.759,34.693L34.665,197.737
		c-46.184,46.186-46.184,121.332,0.001,167.518c22.373,22.373,52.119,34.692,83.758,34.692c31.639,0,61.386-12.319,83.759-34.692
		l163.044-163.044c22.373-22.374,34.694-52.12,34.694-83.759C399.92,86.811,387.599,57.066,365.226,34.693 M191.125,354.197
		c-19.419,19.421-45.239,30.115-72.702,30.115c-27.463,0-53.284-10.694-72.703-30.115c-40.088-40.088-40.088-105.315,0-145.405
		l81.522-81.521c0,0,84.958-84.262,86.215-85.234c1.271-0.953,2.592-1.841,3.879-2.759c5.226-3.572,10.659-6.743,16.269-9.34
		c11.204-5.233,23.011-8.215,34.29-9.053c5.647-0.362,11.15-0.256,16.408,0.298c1.311,0.17,2.608,0.339,3.896,0.506
		c1.281,0.198,2.541,0.455,3.791,0.678c1.251,0.208,2.474,0.502,3.682,0.798c1.204,0.304,2.411,0.54,3.569,0.916
		c1.161,0.353,2.311,0.702,3.447,1.047c1.135,0.339,2.224,0.783,3.315,1.162c2.19,0.744,4.251,1.671,6.264,2.531
		c1.011,0.425,1.965,0.943,2.925,1.4c0.949,0.481,1.906,0.908,2.81,1.414c0.902,0.501,1.791,0.994,2.667,1.479
		c0.436,0.245,0.867,0.488,1.296,0.729c0.417,0.261,0.833,0.521,1.242,0.778c6.596,4.059,11.907,8.118,16.125,11.597
		c5.406,4.459,8.096,6.934,10,9.125c0.953,1.097,2.596,3.101,2.083,3.625c-0.43,0.44-2.131-0.802-4.427-2.563
		c-2.273-1.768-5.523-4.245-9.982-7.139c-4.515-2.813-10.118-5.986-16.807-8.849c-6.666-2.896-14.476-5.357-23.101-6.779
		c-4.32-0.642-8.824-1.161-13.49-1.151c-4.659-0.008-9.453,0.284-14.307,1.044c-9.7,1.536-19.602,4.814-29.036,9.834
		c-4.693,2.555-9.305,5.463-13.732,8.768c-1.089,0.85-81.806,79.059-86.154,83.503c-1.547,1.581-1.463,3.415,0.412,5.29
		c30.383,30.383,115.171,115.08,129.719,129.613c1.39,1.388,1.706,2.559,0.271,4.002
		C245.332,300.136,191.125,354.197,191.125,354.197z"/>
</g>
</svg></div>
      <div class="brand-title">Smart<span>Rx</span></div>
    </div>

    <div class="nav-group">
      <div class="nav-label">Main</div>
      <div id="nav-dispense" class="nav-item active" onclick="openTab('dispense')">
        <span class="nav-icon"><svg fill="currentColor" aria-hidden="true" focusable="false" version="1.1" id="Capa_1" xmlns="http://www.w3.org/2000/svg" 
	 viewBox="0 0 399.947 399.947"
	>
<g>
	<path d="M365.226,34.693C342.854,12.321,313.108,0,281.468,0s-61.386,12.321-83.759,34.693L34.665,197.737
		c-46.184,46.186-46.184,121.332,0.001,167.518c22.373,22.373,52.119,34.692,83.758,34.692c31.639,0,61.386-12.319,83.759-34.692
		l163.044-163.044c22.373-22.374,34.694-52.12,34.694-83.759C399.92,86.811,387.599,57.066,365.226,34.693 M191.125,354.197
		c-19.419,19.421-45.239,30.115-72.702,30.115c-27.463,0-53.284-10.694-72.703-30.115c-40.088-40.088-40.088-105.315,0-145.405
		l81.522-81.521c0,0,84.958-84.262,86.215-85.234c1.271-0.953,2.592-1.841,3.879-2.759c5.226-3.572,10.659-6.743,16.269-9.34
		c11.204-5.233,23.011-8.215,34.29-9.053c5.647-0.362,11.15-0.256,16.408,0.298c1.311,0.17,2.608,0.339,3.896,0.506
		c1.281,0.198,2.541,0.455,3.791,0.678c1.251,0.208,2.474,0.502,3.682,0.798c1.204,0.304,2.411,0.54,3.569,0.916
		c1.161,0.353,2.311,0.702,3.447,1.047c1.135,0.339,2.224,0.783,3.315,1.162c2.19,0.744,4.251,1.671,6.264,2.531
		c1.011,0.425,1.965,0.943,2.925,1.4c0.949,0.481,1.906,0.908,2.81,1.414c0.902,0.501,1.791,0.994,2.667,1.479
		c0.436,0.245,0.867,0.488,1.296,0.729c0.417,0.261,0.833,0.521,1.242,0.778c6.596,4.059,11.907,8.118,16.125,11.597
		c5.406,4.459,8.096,6.934,10,9.125c0.953,1.097,2.596,3.101,2.083,3.625c-0.43,0.44-2.131-0.802-4.427-2.563
		c-2.273-1.768-5.523-4.245-9.982-7.139c-4.515-2.813-10.118-5.986-16.807-8.849c-6.666-2.896-14.476-5.357-23.101-6.779
		c-4.32-0.642-8.824-1.161-13.49-1.151c-4.659-0.008-9.453,0.284-14.307,1.044c-9.7,1.536-19.602,4.814-29.036,9.834
		c-4.693,2.555-9.305,5.463-13.732,8.768c-1.089,0.85-81.806,79.059-86.154,83.503c-1.547,1.581-1.463,3.415,0.412,5.29
		c30.383,30.383,115.171,115.08,129.719,129.613c1.39,1.388,1.706,2.559,0.271,4.002
		C245.332,300.136,191.125,354.197,191.125,354.197z"/>
</g>
</svg></span><span>Dispense</span>
      </div>
      <div id="nav-inventory" class="nav-item" onclick="openTab('inventory')">
        <span class="nav-icon"><svg fill="currentColor" aria-hidden="true" focusable="false" version="1.1" id="Capa_1" xmlns="http://www.w3.org/2000/svg" 
	 viewBox="0 0 450.001 450.001"
	>
<g>
	<path d="M448.554,177.103c0,0-42.15-56.307-42.175-56.339c-0.011-0.015-0.93-1.647-3.762-2.548L226.973,68.62
		c-1.291-0.364-2.655-0.364-3.946,0L47.383,118.216c0,0-2.398,0.692-3.839,2.637c-0.577,0.779-42.097,56.25-42.097,56.25
		c-1.392,1.862-1.81,4.278-1.122,6.499c0.689,2.221,2.4,3.978,4.601,4.725l37.171,12.615v100.874c0,2.938,1.771,5.585,4.487,6.707
		l175.645,72.581c0.522,0.216,1.184,0.551,2.771,0.551s2.771-0.551,2.771-0.551l175.646-72.581c2.716-1.122,4.486-3.77,4.486-6.707
		V200.94l37.171-12.615c2.201-0.747,3.912-2.505,4.601-4.725C450.363,181.381,449.945,178.965,448.554,177.103z M19.062,177.794
		l33.138-44.28l13.403,3.692l75.431,20.778l71.335,19.652l-41.837,51.565L19.062,177.794z M217.743,363.542l-161.13-66.581v-91.094
		l114.04,38.708c0.764,0.261,1.551,0.384,2.332,0.384c2.145,0,4.231-0.953,5.635-2.686l39.124-48.219L217.743,363.542
		L217.743,363.542z M225,166.062L76.37,125.116L225,83.146l148.63,41.97L225,166.062z M393.388,296.961l-161.131,66.581V194.056
		l39.124,48.219c1.404,1.731,3.49,2.686,5.635,2.686c0.781,0,1.568-0.123,2.333-0.384l114.039-38.708V296.961L393.388,296.961z
		 M279.469,229.203l-41.837-51.565l160.169-44.124l33.139,44.28L279.469,229.203z"/>
</g>
</svg></span><span>Inventory</span>
      </div>
      <div id="nav-database" class="nav-item" onclick="openTab('database')">
        <span class="nav-icon"><svg fill="currentColor" aria-hidden="true" focusable="false" id="Layer_1" data-name="Layer 1" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 121.06 122.88"><defs><style>.cls-1{fill-rule:evenodd;}</style></defs><title>database-server</title><path class="cls-1" d="M18,27.16c8.25,2.76,19.74,4.46,32.5,4.46S74.75,29.92,83,27.16c7.24-2.42,11.72-5.35,11.72-8.21S90.24,13.16,83,10.74C74.75,8,63.26,6.28,50.51,6.28S26.26,8,18,10.74C2.65,15.87,2.22,21.89,18,27.16Zm103.05,94c0-.14,0-.26,0-.39V51.37c0-.13,0-.27,0-.39a2.14,2.14,0,0,0-.59-1.11,2.12,2.12,0,0,0-1.5-.63H78.9a2.15,2.15,0,0,0-2.13,2.13v69.37a2.08,2.08,0,0,0,.63,1.5,2,2,0,0,0,1.5.63H119a2.12,2.12,0,0,0,1.5-.63,2.28,2.28,0,0,0,.58-1.12Zm-16.47-15a6.34,6.34,0,0,1-12.41.71l6.78-4.2,5.63,3.49ZM113.2,61.4v5.33H84.67V61.4Zm0,14.74v5.33H99.14L99,81l-.18.44H84.67V76.14ZM6.35,48.73c.57,2.67,4.92,5.36,11.66,7.61,8.25,2.76,19.74,4.46,32.5,4.46,4.45,0,8.75-.21,12.81-.6v6.32c-4.08.38-8.38.58-12.8.58-13.4,0-25.59-1.83-34.48-4.8a39.61,39.61,0,0,1-9.67-4.56V76.57c.57,2.67,4.92,5.36,11.66,7.61,8.25,2.75,19.74,4.46,32.5,4.46,4.44,0,8.74-.21,12.79-.59v6.29c-4.09.37-8.39.57-12.81.57-13.4,0-25.59-1.83-34.48-4.81a39.61,39.61,0,0,1-9.67-4.55v18.83c.57,2.67,4.92,5.36,11.66,7.61,8.25,2.75,19.74,4.46,32.5,4.46,4.45,0,8.75-.21,12.8-.6v4.89c0,.47,0,.94.07,1.41-4.11.38-8.43.58-12.87.58-13.4,0-25.59-1.83-34.48-4.81-4.69-1.56-16-6.7-16-12.85v-2Q0,61,0,19C0,13.08,6.12,8.11,16,4.8,24.91,1.83,37.11,0,50.51,0S76.09,1.83,85,4.81s14.89,7.4,15.88,12.59a3.06,3.06,0,0,1,.18,1.06V35.79H94.75V28.48A39.71,39.71,0,0,1,85,33.09c-8.89,3-21.08,4.81-34.47,4.81S24.92,36.07,16,33.09a39.61,39.61,0,0,1-9.67-4.55V48.73Z"/></svg></span><span>Database</span>
      </div>
      <div id="nav-status" class="nav-item" onclick="openTab('status')">
        <span class="nav-icon"><svg fill="currentColor" aria-hidden="true" focusable="false" version="1.1" id="Capa_1" xmlns="http://www.w3.org/2000/svg" 
	 viewBox="0 0 489.801 489.801">
<g>
	<path d="M484.586,233.212c-5.084-4.578-125.779-111.967-239.689-111.967c-113.903,0-234.598,107.389-239.688,111.967
		C1.891,236.194,0,240.452,0,244.909c0,4.455,1.891,8.711,5.208,11.693c5.09,4.58,125.785,111.953,239.688,111.953
		c113.911,0,234.606-107.373,239.689-111.953c3.322-2.982,5.215-7.238,5.215-11.693
		C489.801,240.452,487.908,236.194,484.586,233.212z M244.896,333.915c-49.068,0-89.005-39.924-89.005-89.006
		c0-49.086,39.937-89.01,89.005-89.01c49.072,0,89.011,39.924,89.011,89.01C333.908,293.991,293.969,333.915,244.896,333.915z
		 M144.137,179.042c-12.422,18.949-19.721,41.553-19.721,65.867c0,24.277,7.268,46.855,19.686,65.803
		c-45.057-21.207-83.982-49.943-103.822-65.818C60.119,229.003,99.045,200.235,144.137,179.042z M345.664,310.757
		c12.416-18.947,19.719-41.551,19.719-65.848c0-24.281-7.285-46.873-19.688-65.822c45.043,21.209,83.969,49.945,103.821,65.822
		C429.677,260.8,390.756,289.55,345.664,310.757z"/>
	<path d="M244.896,194.259c-27.982,0-50.648,22.664-50.648,50.65c0,27.982,22.666,50.648,50.648,50.648
		c27.988,0,50.656-22.666,50.656-50.648C295.553,216.923,272.885,194.259,244.896,194.259z M244.896,226.298
		c-10.262,0-18.607,8.346-18.607,18.611c0,4.348-3.533,7.865-7.867,7.865c-4.365,0-7.869-3.518-7.869-7.865
		c0-18.934,15.4-34.348,34.344-34.348c4.352,0,7.871,3.52,7.871,7.867C252.768,222.78,249.248,226.298,244.896,226.298z"/>
</g>
</svg></span><span>Status</span>
      </div>
    </div>

    <div class="nav-group">
      <div class="nav-label">Settings</div>
      <div id="nav-settings" class="nav-item" onclick="openTab('settings')">
        <span class="nav-icon"><svg fill="currentColor" aria-hidden="true" focusable="false" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg"><path d="M582.239997 663.142968c-56.713997 0-113.426994-21.589999-156.609992-64.772997-11.728999-11.726999-11.728999-30.792998 0-42.520998 11.728999-11.726999 30.791998-11.726999 42.520998 0 62.910997 62.909997 165.212992 62.909997 228.121989 0L914.94998 337.169984c62.910997-62.909997 62.910997-165.211992 0-228.121989-62.907997-62.909997-165.210992-62.909997-228.119989 0L502.313 293.565986c-11.728999 11.727999-30.792998 11.727999-42.521998 0-11.726999-11.727999-11.726999-30.792998 0-42.520998L644.308994 66.527997c86.364996-86.364996 226.858989-86.364996 313.163985 0 86.304996 86.363996 86.363996 226.857989 0 313.163985L738.794989 598.368971c-43.183998 43.183998-99.898995 64.772997-156.612992 64.772997L582.239997 663.141968zM221.386014 1023.99995c-56.714997 0-113.429994-21.591999-156.611992-64.772997-86.363996-86.364996-86.363996-226.856989 0-313.163985l218.678989-218.678989c86.364996-86.364996 226.858989-86.364996 313.162985 0 11.727999 11.727999 11.727999 30.791998 0 42.520998-11.727999 11.726999-30.792998 11.726999-42.518998 0-62.910997-62.909997-165.213992-62.909997-228.122989 0L107.29402 688.583966c-62.909997 62.907997-62.909997 165.212992 0 228.120989 62.908997 62.907997 165.211992 62.907997 228.121989 0l184.516991-184.517991c11.728999-11.728999 30.793998-11.728999 42.520998 0 11.728999 11.726999 11.728999 30.790998 0 42.520998L377.938007 959.226953C334.755009 1002.407951 278.040011 1023.99995 221.325014 1023.99995L221.386014 1023.99995z"  /></svg></span><span>Connection</span>
      </div>
    </div>
  </aside>

  <div class="main">
    <header class="topbar">
      <div class="search">
        <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.25" stroke-linecap="round" stroke-linejoin="round" aria-hidden="true">
          <circle cx="11" cy="11" r="8"></circle>
          <path d="M21 21l-4.35-4.35"></path>
        </svg>
        <span>Search medications, patients…</span>
      </div>

      <div class="topbar-right">
        <span id="badge-edge" class="badge"><span class="dot"></span>Edge: unknown</span>
        <span id="badge-central" class="badge"><span class="dot"></span>Central: unknown</span>
        <span id="badge-mode" class="badge">Mode: setup</span>
        <div class="user-pill">
          <div class="avatar">Rx</div>
          <div class="user-meta">
            <div class="user-name">Pharmacist</div>
            <div class="user-role">Admin</div>
          </div>
        </div>
      </div>
    </header>

    <div class="content">
      <div id="banner-area" class="banner"></div>

      <section id="panel-dispense" class="panel active">
        <div class="page-head">
          <div>
            <div class="title">Dispense<br>Medication</div>
            <div class="sub">Edge first, central fallback with branch override when the edge is down.</div>
          </div>
          <div class="actions">
            <button class="btn ghost" onclick="refreshHealth(true)"><svg fill="currentColor" aria-hidden="true" focusable="false" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg"><path d="M582.239997 663.142968c-56.713997 0-113.426994-21.589999-156.609992-64.772997-11.728999-11.726999-11.728999-30.792998 0-42.520998 11.728999-11.726999 30.791998-11.726999 42.520998 0 62.910997 62.909997 165.212992 62.909997 228.121989 0L914.94998 337.169984c62.910997-62.909997 62.910997-165.211992 0-228.121989-62.907997-62.909997-165.210992-62.909997-228.119989 0L502.313 293.565986c-11.728999 11.727999-30.792998 11.727999-42.521998 0-11.726999-11.727999-11.726999-30.792998 0-42.520998L644.308994 66.527997c86.364996-86.364996 226.858989-86.364996 313.163985 0 86.304996 86.363996 86.363996 226.857989 0 313.163985L738.794989 598.368971c-43.183998 43.183998-99.898995 64.772997-156.612992 64.772997L582.239997 663.141968zM221.386014 1023.99995c-56.714997 0-113.429994-21.591999-156.611992-64.772997-86.363996-86.364996-86.363996-226.856989 0-313.163985l218.678989-218.678989c86.364996-86.364996 226.858989-86.364996 313.162985 0 11.727999 11.727999 11.727999 30.791998 0 42.520998-11.727999 11.726999-30.792998 11.726999-42.518998 0-62.910997-62.909997-165.213992-62.909997-228.122989 0L107.29402 688.583966c-62.909997 62.907997-62.909997 165.212992 0 228.120989 62.908997 62.907997 165.211992 62.907997 228.121989 0l184.516991-184.517991c11.728999-11.728999 30.793998-11.728999 42.520998 0 11.728999 11.726999 11.728999 30.790998 0 42.520998L377.938007 959.226953C334.755009 1002.407951 278.040011 1023.99995 221.325014 1023.99995L221.386014 1023.99995z"  /></svg> Refresh</button>
            <button class="btn primary" onclick="submitDispense()"><svg fill="currentColor" aria-hidden="true" focusable="false" version="1.1" id="Capa_1" xmlns="http://www.w3.org/2000/svg" 
	 viewBox="0 0 399.947 399.947"
	>
<g>
	<path d="M365.226,34.693C342.854,12.321,313.108,0,281.468,0s-61.386,12.321-83.759,34.693L34.665,197.737
		c-46.184,46.186-46.184,121.332,0.001,167.518c22.373,22.373,52.119,34.692,83.758,34.692c31.639,0,61.386-12.319,83.759-34.692
		l163.044-163.044c22.373-22.374,34.694-52.12,34.694-83.759C399.92,86.811,387.599,57.066,365.226,34.693 M191.125,354.197
		c-19.419,19.421-45.239,30.115-72.702,30.115c-27.463,0-53.284-10.694-72.703-30.115c-40.088-40.088-40.088-105.315,0-145.405
		l81.522-81.521c0,0,84.958-84.262,86.215-85.234c1.271-0.953,2.592-1.841,3.879-2.759c5.226-3.572,10.659-6.743,16.269-9.34
		c11.204-5.233,23.011-8.215,34.29-9.053c5.647-0.362,11.15-0.256,16.408,0.298c1.311,0.17,2.608,0.339,3.896,0.506
		c1.281,0.198,2.541,0.455,3.791,0.678c1.251,0.208,2.474,0.502,3.682,0.798c1.204,0.304,2.411,0.54,3.569,0.916
		c1.161,0.353,2.311,0.702,3.447,1.047c1.135,0.339,2.224,0.783,3.315,1.162c2.19,0.744,4.251,1.671,6.264,2.531
		c1.011,0.425,1.965,0.943,2.925,1.4c0.949,0.481,1.906,0.908,2.81,1.414c0.902,0.501,1.791,0.994,2.667,1.479
		c0.436,0.245,0.867,0.488,1.296,0.729c0.417,0.261,0.833,0.521,1.242,0.778c6.596,4.059,11.907,8.118,16.125,11.597
		c5.406,4.459,8.096,6.934,10,9.125c0.953,1.097,2.596,3.101,2.083,3.625c-0.43,0.44-2.131-0.802-4.427-2.563
		c-2.273-1.768-5.523-4.245-9.982-7.139c-4.515-2.813-10.118-5.986-16.807-8.849c-6.666-2.896-14.476-5.357-23.101-6.779
		c-4.32-0.642-8.824-1.161-13.49-1.151c-4.659-0.008-9.453,0.284-14.307,1.044c-9.7,1.536-19.602,4.814-29.036,9.834
		c-4.693,2.555-9.305,5.463-13.732,8.768c-1.089,0.85-81.806,79.059-86.154,83.503c-1.547,1.581-1.463,3.415,0.412,5.29
		c30.383,30.383,115.171,115.08,129.719,129.613c1.39,1.388,1.706,2.559,0.271,4.002
		C245.332,300.136,191.125,354.197,191.125,354.197z"/>
</g>
</svg> Record</button>
          </div>
        </div>

        <div class="grid-4">
          <div class="stat">
            <div class="stat-top">
              <div class="stat-icon"><svg fill="currentColor" aria-hidden="true" focusable="false" version="1.1" id="Capa_1" xmlns="http://www.w3.org/2000/svg" 
	 viewBox="0 0 399.947 399.947"
	>
<g>
	<path d="M365.226,34.693C342.854,12.321,313.108,0,281.468,0s-61.386,12.321-83.759,34.693L34.665,197.737
		c-46.184,46.186-46.184,121.332,0.001,167.518c22.373,22.373,52.119,34.692,83.758,34.692c31.639,0,61.386-12.319,83.759-34.692
		l163.044-163.044c22.373-22.374,34.694-52.12,34.694-83.759C399.92,86.811,387.599,57.066,365.226,34.693 M191.125,354.197
		c-19.419,19.421-45.239,30.115-72.702,30.115c-27.463,0-53.284-10.694-72.703-30.115c-40.088-40.088-40.088-105.315,0-145.405
		l81.522-81.521c0,0,84.958-84.262,86.215-85.234c1.271-0.953,2.592-1.841,3.879-2.759c5.226-3.572,10.659-6.743,16.269-9.34
		c11.204-5.233,23.011-8.215,34.29-9.053c5.647-0.362,11.15-0.256,16.408,0.298c1.311,0.17,2.608,0.339,3.896,0.506
		c1.281,0.198,2.541,0.455,3.791,0.678c1.251,0.208,2.474,0.502,3.682,0.798c1.204,0.304,2.411,0.54,3.569,0.916
		c1.161,0.353,2.311,0.702,3.447,1.047c1.135,0.339,2.224,0.783,3.315,1.162c2.19,0.744,4.251,1.671,6.264,2.531
		c1.011,0.425,1.965,0.943,2.925,1.4c0.949,0.481,1.906,0.908,2.81,1.414c0.902,0.501,1.791,0.994,2.667,1.479
		c0.436,0.245,0.867,0.488,1.296,0.729c0.417,0.261,0.833,0.521,1.242,0.778c6.596,4.059,11.907,8.118,16.125,11.597
		c5.406,4.459,8.096,6.934,10,9.125c0.953,1.097,2.596,3.101,2.083,3.625c-0.43,0.44-2.131-0.802-4.427-2.563
		c-2.273-1.768-5.523-4.245-9.982-7.139c-4.515-2.813-10.118-5.986-16.807-8.849c-6.666-2.896-14.476-5.357-23.101-6.779
		c-4.32-0.642-8.824-1.161-13.49-1.151c-4.659-0.008-9.453,0.284-14.307,1.044c-9.7,1.536-19.602,4.814-29.036,9.834
		c-4.693,2.555-9.305,5.463-13.732,8.768c-1.089,0.85-81.806,79.059-86.154,83.503c-1.547,1.581-1.463,3.415,0.412,5.29
		c30.383,30.383,115.171,115.08,129.719,129.613c1.39,1.388,1.706,2.559,0.271,4.002
		C245.332,300.136,191.125,354.197,191.125,354.197z"/>
</g>
</svg></div>
              <span id="kpi-dispense-chip" class="stat-chip warn">checking</span>
            </div>
            <div class="stat-label">Dispense path</div>
            <div id="kpi-dispense" class="stat-value">—</div>
          </div>
          <div class="stat">
            <div class="stat-top">
              <div class="stat-icon"><svg fill="currentColor" aria-hidden="true" focusable="false" id="Layer_1" data-name="Layer 1" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 121.06 122.88"><defs><style>.cls-1{fill-rule:evenodd;}</style></defs><title>database-server</title><path class="cls-1" d="M18,27.16c8.25,2.76,19.74,4.46,32.5,4.46S74.75,29.92,83,27.16c7.24-2.42,11.72-5.35,11.72-8.21S90.24,13.16,83,10.74C74.75,8,63.26,6.28,50.51,6.28S26.26,8,18,10.74C2.65,15.87,2.22,21.89,18,27.16Zm103.05,94c0-.14,0-.26,0-.39V51.37c0-.13,0-.27,0-.39a2.14,2.14,0,0,0-.59-1.11,2.12,2.12,0,0,0-1.5-.63H78.9a2.15,2.15,0,0,0-2.13,2.13v69.37a2.08,2.08,0,0,0,.63,1.5,2,2,0,0,0,1.5.63H119a2.12,2.12,0,0,0,1.5-.63,2.28,2.28,0,0,0,.58-1.12Zm-16.47-15a6.34,6.34,0,0,1-12.41.71l6.78-4.2,5.63,3.49ZM113.2,61.4v5.33H84.67V61.4Zm0,14.74v5.33H99.14L99,81l-.18.44H84.67V76.14ZM6.35,48.73c.57,2.67,4.92,5.36,11.66,7.61,8.25,2.76,19.74,4.46,32.5,4.46,4.45,0,8.75-.21,12.81-.6v6.32c-4.08.38-8.38.58-12.8.58-13.4,0-25.59-1.83-34.48-4.8a39.61,39.61,0,0,1-9.67-4.56V76.57c.57,2.67,4.92,5.36,11.66,7.61,8.25,2.75,19.74,4.46,32.5,4.46,4.44,0,8.74-.21,12.79-.59v6.29c-4.09.37-8.39.57-12.81.57-13.4,0-25.59-1.83-34.48-4.81a39.61,39.61,0,0,1-9.67-4.55v18.83c.57,2.67,4.92,5.36,11.66,7.61,8.25,2.75,19.74,4.46,32.5,4.46,4.45,0,8.75-.21,12.8-.6v4.89c0,.47,0,.94.07,1.41-4.11.38-8.43.58-12.87.58-13.4,0-25.59-1.83-34.48-4.81-4.69-1.56-16-6.7-16-12.85v-2Q0,61,0,19C0,13.08,6.12,8.11,16,4.8,24.91,1.83,37.11,0,50.51,0S76.09,1.83,85,4.81s14.89,7.4,15.88,12.59a3.06,3.06,0,0,1,.18,1.06V35.79H94.75V28.48A39.71,39.71,0,0,1,85,33.09c-8.89,3-21.08,4.81-34.47,4.81S24.92,36.07,16,33.09a39.61,39.61,0,0,1-9.67-4.55V48.73Z"/></svg></div>
              <span id="kpi-db-chip" class="stat-chip warn">checking</span>
            </div>
            <div class="stat-label">Database view</div>
            <div id="kpi-database" class="stat-value">—</div>
          </div>
          <div class="stat">
            <div class="stat-top">
              <div class="stat-icon"><svg fill="currentColor" aria-hidden="true" focusable="false" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg"><path d="M582.239997 663.142968c-56.713997 0-113.426994-21.589999-156.609992-64.772997-11.728999-11.726999-11.728999-30.792998 0-42.520998 11.728999-11.726999 30.791998-11.726999 42.520998 0 62.910997 62.909997 165.212992 62.909997 228.121989 0L914.94998 337.169984c62.910997-62.909997 62.910997-165.211992 0-228.121989-62.907997-62.909997-165.210992-62.909997-228.119989 0L502.313 293.565986c-11.728999 11.727999-30.792998 11.727999-42.521998 0-11.726999-11.727999-11.726999-30.792998 0-42.520998L644.308994 66.527997c86.364996-86.364996 226.858989-86.364996 313.163985 0 86.304996 86.363996 86.363996 226.857989 0 313.163985L738.794989 598.368971c-43.183998 43.183998-99.898995 64.772997-156.612992 64.772997L582.239997 663.141968zM221.386014 1023.99995c-56.714997 0-113.429994-21.591999-156.611992-64.772997-86.363996-86.364996-86.363996-226.856989 0-313.163985l218.678989-218.678989c86.364996-86.364996 226.858989-86.364996 313.162985 0 11.727999 11.727999 11.727999 30.791998 0 42.520998-11.727999 11.726999-30.792998 11.726999-42.518998 0-62.910997-62.909997-165.213992-62.909997-228.122989 0L107.29402 688.583966c-62.909997 62.907997-62.909997 165.212992 0 228.120989 62.908997 62.907997 165.211992 62.907997 228.121989 0l184.516991-184.517991c11.728999-11.728999 30.793998-11.728999 42.520998 0 11.728999 11.726999 11.728999 30.790998 0 42.520998L377.938007 959.226953C334.755009 1002.407951 278.040011 1023.99995 221.325014 1023.99995L221.386014 1023.99995z"  /></svg></div>
              <span class="stat-chip neutral">branch</span>
            </div>
            <div class="stat-label">Branch ID</div>
            <div id="kpi-branch" class="stat-value">—</div>
          </div>
          <div class="stat">
            <div class="stat-top">
              <div class="stat-icon"><svg fill="currentColor" aria-hidden="true" focusable="false" version="1.1" id="Capa_1" xmlns="http://www.w3.org/2000/svg" 
	 viewBox="0 0 489.801 489.801">
<g>
	<path d="M484.586,233.212c-5.084-4.578-125.779-111.967-239.689-111.967c-113.903,0-234.598,107.389-239.688,111.967
		C1.891,236.194,0,240.452,0,244.909c0,4.455,1.891,8.711,5.208,11.693c5.09,4.58,125.785,111.953,239.688,111.953
		c113.911,0,234.606-107.373,239.689-111.953c3.322-2.982,5.215-7.238,5.215-11.693
		C489.801,240.452,487.908,236.194,484.586,233.212z M244.896,333.915c-49.068,0-89.005-39.924-89.005-89.006
		c0-49.086,39.937-89.01,89.005-89.01c49.072,0,89.011,39.924,89.011,89.01C333.908,293.991,293.969,333.915,244.896,333.915z
		 M144.137,179.042c-12.422,18.949-19.721,41.553-19.721,65.867c0,24.277,7.268,46.855,19.686,65.803
		c-45.057-21.207-83.982-49.943-103.822-65.818C60.119,229.003,99.045,200.235,144.137,179.042z M345.664,310.757
		c12.416-18.947,19.719-41.551,19.719-65.848c0-24.281-7.285-46.873-19.688-65.822c45.043,21.209,83.969,49.945,103.821,65.822
		C429.677,260.8,390.756,289.55,345.664,310.757z"/>
	<path d="M244.896,194.259c-27.982,0-50.648,22.664-50.648,50.65c0,27.982,22.666,50.648,50.648,50.648
		c27.988,0,50.656-22.666,50.656-50.648C295.553,216.923,272.885,194.259,244.896,194.259z M244.896,226.298
		c-10.262,0-18.607,8.346-18.607,18.611c0,4.348-3.533,7.865-7.867,7.865c-4.365,0-7.869-3.518-7.869-7.865
		c0-18.934,15.4-34.348,34.344-34.348c4.352,0,7.871,3.52,7.871,7.867C252.768,222.78,249.248,226.298,244.896,226.298z"/>
</g>
</svg></div>
              <span id="kpi-nodes-chip" class="stat-chip warn">nodes</span>
            </div>
            <div class="stat-label">Reachable nodes</div>
            <div id="kpi-nodes" class="stat-value">0 / 2</div>
          </div>
        </div>

        <div class="two-col">
          <div class="card">
            <div class="card-title">New Dispense</div>
            <div class="form-grid">
              <div class="f2">
                <label for="d-med">Medication</label>
                <select id="d-med"></select>
              </div>
              <div class="f1">
                <label for="d-qty">Quantity</label>
                <input id="d-qty" type="number" min="1" value="1">
              </div>
              <div class="f1">
                <label for="d-pharm">Pharmacist ID</label>
                <input id="d-pharm" placeholder="PH-001">
              </div>
              <div class="f2">
                <label for="d-patient">Patient ID</label>
                <input id="d-patient" placeholder="P-0042">
              </div>
              <div class="f12">
                <button class="btn primary" style="width:100%;" onclick="submitDispense()"><svg fill="currentColor" aria-hidden="true" focusable="false" version="1.1" id="Capa_1" xmlns="http://www.w3.org/2000/svg" 
	 viewBox="0 0 399.947 399.947"
	>
<g>
	<path d="M365.226,34.693C342.854,12.321,313.108,0,281.468,0s-61.386,12.321-83.759,34.693L34.665,197.737
		c-46.184,46.186-46.184,121.332,0.001,167.518c22.373,22.373,52.119,34.692,83.758,34.692c31.639,0,61.386-12.319,83.759-34.692
		l163.044-163.044c22.373-22.374,34.694-52.12,34.694-83.759C399.92,86.811,387.599,57.066,365.226,34.693 M191.125,354.197
		c-19.419,19.421-45.239,30.115-72.702,30.115c-27.463,0-53.284-10.694-72.703-30.115c-40.088-40.088-40.088-105.315,0-145.405
		l81.522-81.521c0,0,84.958-84.262,86.215-85.234c1.271-0.953,2.592-1.841,3.879-2.759c5.226-3.572,10.659-6.743,16.269-9.34
		c11.204-5.233,23.011-8.215,34.29-9.053c5.647-0.362,11.15-0.256,16.408,0.298c1.311,0.17,2.608,0.339,3.896,0.506
		c1.281,0.198,2.541,0.455,3.791,0.678c1.251,0.208,2.474,0.502,3.682,0.798c1.204,0.304,2.411,0.54,3.569,0.916
		c1.161,0.353,2.311,0.702,3.447,1.047c1.135,0.339,2.224,0.783,3.315,1.162c2.19,0.744,4.251,1.671,6.264,2.531
		c1.011,0.425,1.965,0.943,2.925,1.4c0.949,0.481,1.906,0.908,2.81,1.414c0.902,0.501,1.791,0.994,2.667,1.479
		c0.436,0.245,0.867,0.488,1.296,0.729c0.417,0.261,0.833,0.521,1.242,0.778c6.596,4.059,11.907,8.118,16.125,11.597
		c5.406,4.459,8.096,6.934,10,9.125c0.953,1.097,2.596,3.101,2.083,3.625c-0.43,0.44-2.131-0.802-4.427-2.563
		c-2.273-1.768-5.523-4.245-9.982-7.139c-4.515-2.813-10.118-5.986-16.807-8.849c-6.666-2.896-14.476-5.357-23.101-6.779
		c-4.32-0.642-8.824-1.161-13.49-1.151c-4.659-0.008-9.453,0.284-14.307,1.044c-9.7,1.536-19.602,4.814-29.036,9.834
		c-4.693,2.555-9.305,5.463-13.732,8.768c-1.089,0.85-81.806,79.059-86.154,83.503c-1.547,1.581-1.463,3.415,0.412,5.29
		c30.383,30.383,115.171,115.08,129.719,129.613c1.39,1.388,1.706,2.559,0.271,4.002
		C245.332,300.136,191.125,354.197,191.125,354.197z"/>
</g>
</svg> Record Dispense</button>
              </div>
            </div>
            <p class="footer-note">Dispense uses the edge node when available. If the edge is unavailable, the request is sent to central with your branch ID.</p>
          </div>

          <div class="card">
            <div class="card-title">Recent Transactions</div>
            <div id="recent-source" class="footer-note" style="margin-top:0; margin-bottom:0.6rem;">Source: —</div>
            <div id="recent-tx"></div>
          </div>
        </div>
      </section>

      <section id="panel-inventory" class="panel">
        <div class="page-head">
          <div>
            <div class="title">Operational<br>Inventory</div>
            <div class="sub">Reads from the active node for quick branch-level stock visibility.</div>
          </div>
          <div class="actions">
            <button class="btn ghost" onclick="loadInventoryTable()"><svg fill="currentColor" aria-hidden="true" focusable="false" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg"><path d="M582.239997 663.142968c-56.713997 0-113.426994-21.589999-156.609992-64.772997-11.728999-11.726999-11.728999-30.792998 0-42.520998 11.728999-11.726999 30.791998-11.726999 42.520998 0 62.910997 62.909997 165.212992 62.909997 228.121989 0L914.94998 337.169984c62.910997-62.909997 62.910997-165.211992 0-228.121989-62.907997-62.909997-165.210992-62.909997-228.119989 0L502.313 293.565986c-11.728999 11.727999-30.792998 11.727999-42.521998 0-11.726999-11.727999-11.726999-30.792998 0-42.520998L644.308994 66.527997c86.364996-86.364996 226.858989-86.364996 313.163985 0 86.304996 86.363996 86.363996 226.857989 0 313.163985L738.794989 598.368971c-43.183998 43.183998-99.898995 64.772997-156.612992 64.772997L582.239997 663.141968zM221.386014 1023.99995c-56.714997 0-113.429994-21.591999-156.611992-64.772997-86.363996-86.364996-86.363996-226.856989 0-313.163985l218.678989-218.678989c86.364996-86.364996 226.858989-86.364996 313.162985 0 11.727999 11.727999 11.727999 30.791998 0 42.520998-11.727999 11.726999-30.792998 11.726999-42.518998 0-62.910997-62.909997-165.213992-62.909997-228.122989 0L107.29402 688.583966c-62.909997 62.907997-62.909997 165.212992 0 228.120989 62.908997 62.907997 165.211992 62.907997 228.121989 0l184.516991-184.517991c11.728999-11.728999 30.793998-11.728999 42.520998 0 11.728999 11.726999 11.728999 30.790998 0 42.520998L377.938007 959.226953C334.755009 1002.407951 278.040011 1023.99995 221.325014 1023.99995L221.386014 1023.99995z"  /></svg> Refresh</button>
          </div>
        </div>
        <div class="card">
          <div id="inventory-source" class="footer-note" style="margin-top:0; margin-bottom:0.6rem;">Source: —</div>
          <div id="inventory-table"></div>
        </div>
      </section>

      <section id="panel-database" class="panel">
        <div class="page-head">
          <div>
            <div class="title">Database<br>Administration</div>
            <div class="sub">Medications, stock, patients, and transactions — central only.</div>
          </div>
        </div>
        <div id="database-lock" class="banner warn">Central node is unavailable. The database view is disabled.</div>

        <div id="database-panel" class="overlay">
          <div class="two-col" style="margin-bottom:0.9rem;">
            <div class="card">
              <div class="card-title">Register Medication</div>
              <div class="form-grid">
                <div class="f1">
                  <label for="m-id">Medication ID</label>
                  <input id="m-id" placeholder="MED-001">
                </div>
                <div class="f2">
                  <label for="m-name">Name</label>
                  <input id="m-name" placeholder="Paracetamol 500mg">
                </div>
                <div class="f1">
                  <label for="m-unit">Unit</label>
                  <input id="m-unit" placeholder="tablet" value="tablet">
                </div>
                <div class="f1">
                  <label for="m-min">Min stock</label>
                  <input id="m-min" type="number" value="10">
                </div>
                <div class="f12">
                  <button id="btn-med-save" class="btn success" style="width:100%;" onclick="addMedication()">Save Medication</button>
                </div>
              </div>
            </div>

            <div class="card">
              <div class="card-title">Add Stock</div>
              <div class="form-grid">
                <div class="f2">
                  <label for="s-med">Medication</label>
                  <select id="s-med"></select>
                </div>
                <div class="f1">
                  <label for="s-qty">Quantity</label>
                  <input id="s-qty" type="number" min="1" value="10">
                </div>
                <div class="f1">
                  <label for="s-by">Added By</label>
                  <input id="s-by" placeholder="ADM-001">
                </div>
                <div class="f12">
                  <button id="btn-stock-add" class="btn success" style="width:100%;" onclick="addStock()">Add Stock</button>
                </div>
              </div>
            </div>
          </div>

          <div class="two-col">
            <div class="card">
              <div class="card-title">Patient Registry</div>
              <div class="form-grid">
                <div class="f1">
                  <label for="p-id">Patient ID</label>
                  <input id="p-id" placeholder="P-0001">
                </div>
                <div class="f2">
                  <label for="p-name">Name</label>
                  <input id="p-name" placeholder="Ahmed Mohamed">
                </div>
                <div class="f1">
                  <label for="p-dob">Date of birth</label>
                  <input id="p-dob" type="date">
                </div>
                <div class="f2">
                  <label for="p-contact">Contact</label>
                  <input id="p-contact" placeholder="+20 1xx xxx xxxx">
                </div>
                <div class="f1">
                  <label for="p-branch">Branch ID</label>
                  <input id="p-branch" placeholder="branch-1">
                </div>
                <div class="f12">
                  <button id="btn-patient-save" class="btn primary" style="width:100%;" onclick="savePatient()">Save Patient</button>
                </div>
              </div>
            </div>

            <div class="card">
              <div class="card-title">Database Tables</div>
              <div class="tabs">
                <button id="db-tab-medications" class="tab-btn active" onclick="setDatabaseView('medications', this)">Medications</button>
                <button id="db-tab-patients" class="tab-btn" onclick="setDatabaseView('patients', this)">Patients</button>
                <button id="db-tab-transactions" class="tab-btn" onclick="setDatabaseView('transactions', this)">Transactions</button>
              </div>
              <div id="db-view"></div>
            </div>
          </div>
        </div>
      </section>

      <section id="panel-status" class="panel">
        <div class="page-head">
          <div>
            <div class="title">Node<br>Status</div>
            <div class="sub">Live health, replication-aware status, and routing decisions.</div>
          </div>
          <div class="actions">
            <button class="btn ghost" onclick="refreshHealth(true)"><svg fill="currentColor" aria-hidden="true" focusable="false" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg"><path d="M582.239997 663.142968c-56.713997 0-113.426994-21.589999-156.609992-64.772997-11.728999-11.726999-11.728999-30.792998 0-42.520998 11.728999-11.726999 30.791998-11.726999 42.520998 0 62.910997 62.909997 165.212992 62.909997 228.121989 0L914.94998 337.169984c62.910997-62.909997 62.910997-165.211992 0-228.121989-62.907997-62.909997-165.210992-62.909997-228.119989 0L502.313 293.565986c-11.728999 11.727999-30.792998 11.727999-42.521998 0-11.726999-11.727999-11.726999-30.792998 0-42.520998L644.308994 66.527997c86.364996-86.364996 226.858989-86.364996 313.163985 0 86.304996 86.363996 86.363996 226.857989 0 313.163985L738.794989 598.368971c-43.183998 43.183998-99.898995 64.772997-156.612992 64.772997L582.239997 663.141968zM221.386014 1023.99995c-56.714997 0-113.429994-21.591999-156.611992-64.772997-86.363996-86.364996-86.363996-226.856989 0-313.163985l218.678989-218.678989c86.364996-86.364996 226.858989-86.364996 313.162985 0 11.727999 11.727999 11.727999 30.791998 0 42.520998-11.727999 11.726999-30.792998 11.726999-42.518998 0-62.910997-62.909997-165.213992-62.909997-228.122989 0L107.29402 688.583966c-62.909997 62.907997-62.909997 165.212992 0 228.120989 62.908997 62.907997 165.211992 62.907997 228.121989 0l184.516991-184.517991c11.728999-11.728999 30.793998-11.728999 42.520998 0 11.728999 11.726999 11.728999 30.790998 0 42.520998L377.938007 959.226953C334.755009 1002.407951 278.040011 1023.99995 221.325014 1023.99995L221.386014 1023.99995z"  /></svg> Refresh</button>
          </div>
        </div>

        <div class="two-col" style="margin-bottom:0.9rem;">
          <div class="card">
            <div class="card-title">Edge Node</div>
            <div id="edge-chip" class="chip neutral">Checking…</div>
            <div id="edge-details" class="status-line"></div>
          </div>
          <div class="card">
            <div class="card-title">Central Node</div>
            <div id="central-chip" class="chip neutral">Checking…</div>
            <div id="central-details" class="status-line"></div>
          </div>
        </div>

        <div class="two-col">
          <div class="card">
            <div class="card-title">Current Routing</div>
            <div id="routing-summary" class="status-line"></div>
          </div>
          <div class="card">
            <div class="card-title">Replication Note</div>
            <p class="footer-note" style="margin-top:0; font-size:0.88rem; line-height:1.6;">
              Marmot is expected to watch each SQLite WAL file and replicate changes asynchronously. The UI only reacts to node health and routes actions to the correct node.
            </p>
          </div>
        </div>
      </section>

      <section id="panel-settings" class="panel">
        <div class="page-head">
          <div>
            <div class="title">Connection<br>Settings</div>
            <div class="sub">Configure node URLs and branch ID. Saved automatically in your browser.</div>
          </div>
        </div>

        <div class="card" style="max-width: 860px;">
          <div class="card-title">Node Configuration</div>
          <div class="settings-grid">
            <div>
              <label for="edge-api">Edge Node URL</label>
              <input id="edge-api" placeholder="http://127.0.0.1:8080">
            </div>
            <div>
              <label for="central-api">Central Node URL</label>
              <input id="central-api" placeholder="http://127.0.0.1:8080">
            </div>
            <div>
              <label for="branch-id">Branch ID</label>
              <input id="branch-id" placeholder="branch-1">
            </div>
            <div>
              <label for="auth-token">Auth Token</label>
              <input id="auth-token" type="password" placeholder="dev-pharmacist-token">
            </div>
            <div style="display:flex; gap:0.5rem; align-items:flex-end;">
              <button class="btn primary" onclick="saveSettings()">Save</button>
              <button class="btn ghost icon-btn" onclick="refreshHealth(true)"><svg fill="currentColor" aria-hidden="true" focusable="false" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg"><path d="M582.239997 663.142968c-56.713997 0-113.426994-21.589999-156.609992-64.772997-11.728999-11.726999-11.728999-30.792998 0-42.520998 11.728999-11.726999 30.791998-11.726999 42.520998 0 62.910997 62.909997 165.212992 62.909997 228.121989 0L914.94998 337.169984c62.910997-62.909997 62.910997-165.211992 0-228.121989-62.907997-62.909997-165.210992-62.909997-228.119989 0L502.313 293.565986c-11.728999 11.727999-30.792998 11.727999-42.521998 0-11.726999-11.727999-11.726999-30.792998 0-42.520998L644.308994 66.527997c86.364996-86.364996 226.858989-86.364996 313.163985 0 86.304996 86.363996 86.363996 226.857989 0 313.163985L738.794989 598.368971c-43.183998 43.183998-99.898995 64.772997-156.612992 64.772997L582.239997 663.141968zM221.386014 1023.99995c-56.714997 0-113.429994-21.591999-156.611992-64.772997-86.363996-86.364996-86.363996-226.856989 0-313.163985l218.678989-218.678989c86.364996-86.364996 226.858989-86.364996 313.162985 0 11.727999 11.727999 11.727999 30.791998 0 42.520998-11.727999 11.726999-30.792998 11.726999-42.518998 0-62.910997-62.909997-165.213992-62.909997-228.122989 0L107.29402 688.583966c-62.909997 62.907997-62.909997 165.212992 0 228.120989 62.908997 62.907997 165.211992 62.907997 228.121989 0l184.516991-184.517991c11.728999-11.728999 30.793998-11.728999 42.520998 0 11.728999 11.726999 11.728999 30.790998 0 42.520998L377.938007 959.226953C334.755009 1002.407951 278.040011 1023.99995 221.325014 1023.99995L221.386014 1023.99995z"  /></svg></button>
            </div>
          </div>
          <p class="footer-note">You can edit these values once. They persist in localStorage and also update live when changed.</p>
        </div>

        <div class="two-col" style="max-width: 860px; margin-top: 1rem;">
          <div class="card">
            <div class="card-title">Edge Node</div>
            <div id="edge-chip-2" class="chip neutral">Checking…</div>
            <div id="edge-details-2" class="status-line"></div>
          </div>
          <div class="card">
            <div class="card-title">Central Node</div>
            <div id="central-chip-2" class="chip neutral">Checking…</div>
            <div id="central-details-2" class="status-line"></div>
          </div>
        </div>
      </section>
    </div>
  </div>
</div>

<script>
const STORAGE = {
  edge: 'smart-pharmacy.edge_api',
  central: 'smart-pharmacy.central_api',
  branch: 'smart-pharmacy.branch_id',
  auth: 'smart-pharmacy.auth_token'
};
const DEFAULT_AUTH_TOKEN = 'dev-pharmacist-token';

let EDGE_API = localStorage.getItem(STORAGE.edge) || '';
let CENTRAL_API = localStorage.getItem(STORAGE.central) || '';
let BRANCH_ID = localStorage.getItem(STORAGE.branch) || '';
let AUTH_TOKEN = normalizeAuthToken(localStorage.getItem(STORAGE.auth));

let edgeState = { alive: false, info: null, error: 'not checked yet' };
let centralState = { alive: false, info: null, error: 'not checked yet' };
let activeTab = 'dispense';
let dbView = 'medications';
let pollTimer = null;
let edgeSeedAttempted = false;

const STATUS_PATHS = ['/api/status', '/status', '/health', '/api/health'];

function esc(v) {
  return String(v == null ? '' : v);
}

function setBanner(type, msg) {
  const el = document.getElementById('banner-area');
  if (!msg) {
    el.style.display = 'none';
    el.textContent = '';
    return;
  }
  el.className = 'banner ' + type;
  el.textContent = msg;
  el.style.display = 'block';
}

function makeChip(text, kind) {
  const span = document.createElement('span');
  span.className = 'chip ' + (kind || 'neutral');
  span.textContent = text;
  return span;
}

function setStatusChip(el, alive, label) {
  el.className = 'chip ' + (alive ? 'good' : 'bad');
  el.textContent = '● ' + label;
}

function setStatChip(el, kind, text) {
  el.className = 'stat-chip ' + kind;
  el.textContent = text;
}

function currentInput(id) {
  const el = document.getElementById(id);
  return el ? el.value.trim() : '';
}

function normalizeAuthToken(value) {
  return String(value || '').trim() || DEFAULT_AUTH_TOKEN;
}

function persistAuthToken(value) {
  AUTH_TOKEN = normalizeAuthToken(value);
  localStorage.setItem(STORAGE.auth, AUTH_TOKEN);
  const authEl = document.getElementById('auth-token');
  if (authEl && document.activeElement !== authEl) {
    authEl.value = AUTH_TOKEN;
  }
}

function loadSettingsFromStorage() {
  const edgeEl = document.getElementById('edge-api');
  const centralEl = document.getElementById('central-api');
  const branchEl = document.getElementById('branch-id');
  const authEl = document.getElementById('auth-token');
  if (edgeEl) edgeEl.value = EDGE_API;
  if (centralEl) centralEl.value = CENTRAL_API;
  if (branchEl) branchEl.value = BRANCH_ID;
  if (authEl) authEl.value = AUTH_TOKEN;
}

function syncSettingsFromInputs() {
  EDGE_API = currentInput('edge-api');
  CENTRAL_API = currentInput('central-api');
  BRANCH_ID = currentInput('branch-id');
  AUTH_TOKEN = normalizeAuthToken(currentInput('auth-token'));
  localStorage.setItem(STORAGE.edge, EDGE_API);
  localStorage.setItem(STORAGE.central, CENTRAL_API);
  localStorage.setItem(STORAGE.branch, BRANCH_ID);
  localStorage.setItem(STORAGE.auth, AUTH_TOKEN);
  updateInputsLive();
}

function updateInputsLive() {
  const fields = [
    ['edge-api', EDGE_API],
    ['central-api', CENTRAL_API],
    ['branch-id', BRANCH_ID],
    ['auth-token', AUTH_TOKEN],
  ];
  fields.forEach(([id, value]) => {
    const el = document.getElementById(id);
    if (!el) return;
    if (document.activeElement !== el && el.value !== value) {
      el.value = value;
    }
  });
}

async function fetchJson(base, path, opts, timeoutMs) {
  const timeout = typeof timeoutMs === 'number' ? timeoutMs : 2200;
  if (!base) throw new Error('URL not set');
  const url = base.replace(/\/+$/, '') + path;
  const controller = new AbortController();
  const timer = setTimeout(() => controller.abort(), timeout);
  const requestOpts = opts || {};
  const headers = new Headers(requestOpts.headers || {});
  if (!headers.has('Authorization') && !headers.has('X-Pharmacy-Token')) {
    headers.set('Authorization', 'Bearer ' + normalizeAuthToken(AUTH_TOKEN));
  }

  try {
    const response = await fetch(url, {
      ...requestOpts,
      headers,
      signal: controller.signal,
    });
    const parsed = await parseJsonResponse(response);
    if (parsed.ok) return parsed.data;

    const isAuthFailure = response.status === 401 && String(parsed.message).toLowerCase().includes('authentication');
    const canRetryWithDefault = isAuthFailure && AUTH_TOKEN !== DEFAULT_AUTH_TOKEN &&
      !new Headers(requestOpts.headers || {}).has('Authorization') &&
      !new Headers(requestOpts.headers || {}).has('X-Pharmacy-Token');
    if (canRetryWithDefault) {
      persistAuthToken(DEFAULT_AUTH_TOKEN);
      const retryHeaders = new Headers(requestOpts.headers || {});
      retryHeaders.set('Authorization', 'Bearer ' + DEFAULT_AUTH_TOKEN);
      const retryRes = await fetch(url, {
        ...requestOpts,
        headers: retryHeaders,
        signal: controller.signal,
      });
      const retryParsed = await parseJsonResponse(retryRes);
      if (retryParsed.ok) return retryParsed.data;
      throw httpError(retryParsed.message, retryRes.status, retryParsed.data);
    }

    if (isAuthFailure) {
      throw httpError('Authentication required. Check the Auth Token in Settings.', response.status, parsed.data);
    }
    throw httpError(parsed.message, response.status, parsed.data);
  } finally {
    clearTimeout(timer);
  }
}

async function parseJsonResponse(res) {
    const text = await res.text();
    let data = null;
    try {
      data = text ? JSON.parse(text) : null;
    } catch {
      data = text;
    }
    const message = data && data.error ? data.error : ('request failed (' + res.status + ')');
    return { ok: res.ok, data: data, message: message };
}

function httpError(message, status, data) {
  const err = new Error(message);
  err.status = status;
  err.data = data;
  return err;
}

function newClientID(prefix) {
  if (window.crypto && crypto.randomUUID) {
    return prefix + '-' + crypto.randomUUID();
  }
  return prefix + '-' + Date.now() + '-' + Math.random().toString(36).slice(2);
}

function sqlNowUTC() {
  return new Date().toISOString().replace('T', ' ').replace(/\.\d{3}Z$/, '');
}

async function probeStatus(base) {
  if (!base) return { alive: false, info: null, error: 'URL not set' };

  for (const path of STATUS_PATHS) {
    try {
      const info = await fetchJson(base, path, {}, 1800);
      return { alive: true, info: info, error: '' };
    } catch (err) {
    }
  }
  return { alive: false, info: null, error: 'offline or unreachable' };
}

function statusBadges(info) {
  const out = [];
  if (!info) return out;
  if (info.node_id) out.push(['node ' + info.node_id, 'neutral']);
  if (info.role) out.push([info.role, 'good']);
  if (typeof info.branch !== 'undefined' && info.branch !== '') out.push(['branch ' + info.branch, 'neutral']);
  if (info.replication_state) out.push(['replication ' + info.replication_state, 'neutral']);
  if (Array.isArray(info.capabilities)) out.push(['capabilities ' + info.capabilities.join(', '), 'neutral']);
  return out;
}

function updateNodeCard(prefix, state, base) {
  const chip = document.getElementById(prefix + '-chip');
  const chip2 = document.getElementById(prefix + '-chip-2');
  const details = document.getElementById(prefix + '-details');
  const details2 = document.getElementById(prefix + '-details-2');

  const onlineLabel = state.alive ? 'online' : (base ? 'offline' : 'URL not set');
  setStatusChip(chip, state.alive, onlineLabel);
  if (chip2) setStatusChip(chip2, state.alive, onlineLabel);

  const fill = (box) => {
    if (!box) return;
    box.innerHTML = '';
    const badges = state.alive ? statusBadges(state.info) : [
      [state.error || onlineLabel, state.alive ? 'good' : 'bad']
    ];
    if (badges.length === 0) badges.push(['no details', 'neutral']);
    badges.forEach(([text, kind]) => box.appendChild(makeChip(text, kind)));
  };

  fill(details);
  fill(details2);
}

function updateUiState() {
  const edgeAlive = edgeState.alive;
  const centralAlive = centralState.alive;
  const activeNodes = (edgeAlive ? 1 : 0) + (centralAlive ? 1 : 0);

  const mode = document.getElementById('badge-mode');
  if (edgeAlive && centralAlive) {
    mode.className = 'badge good';
    mode.innerHTML = '<span class="dot"></span>Fully available';
  } else if (edgeAlive && !centralAlive) {
    mode.className = 'badge warn';
    mode.innerHTML = '<span class="dot"></span>Edge only';
  } else if (!edgeAlive && centralAlive) {
    mode.className = 'badge warn';
    mode.innerHTML = '<span class="dot"></span>Central fallback';
  } else {
    mode.className = 'badge bad';
    mode.innerHTML = '<span class="dot"></span>No nodes';
  }

  const be = document.getElementById('badge-edge');
  be.className = 'badge ' + (edgeAlive ? 'good' : 'bad');
  be.innerHTML = '<span class="dot"></span>Edge: ' + (edgeAlive ? 'online' : 'offline');

  const bc = document.getElementById('badge-central');
  bc.className = 'badge ' + (centralAlive ? 'good' : 'bad');
  bc.innerHTML = '<span class="dot"></span>Central: ' + (centralAlive ? 'online' : 'offline');

  document.getElementById('kpi-dispense').textContent = edgeAlive ? 'Edge' : (centralAlive ? 'Central' : '—');
  document.getElementById('kpi-database').textContent = centralAlive ? 'Central' : 'Disabled';
  document.getElementById('kpi-branch').textContent = BRANCH_ID || 'not set';
  document.getElementById('kpi-nodes').textContent = activeNodes + ' / 2';

  setStatChip(document.getElementById('kpi-dispense-chip'), edgeAlive ? 'up' : (centralAlive ? 'warn' : 'down'), edgeAlive ? 'edge' : (centralAlive ? 'fallback' : 'down'));
  setStatChip(document.getElementById('kpi-db-chip'), centralAlive ? 'up' : 'down', centralAlive ? 'live' : 'offline');
  setStatChip(document.getElementById('kpi-nodes-chip'), activeNodes === 2 ? 'up' : (activeNodes === 1 ? 'warn' : 'down'), activeNodes === 2 ? 'healthy' : (activeNodes === 1 ? 'partial' : 'down'));

  document.getElementById('nav-database').classList.toggle('locked', !centralAlive);
  document.getElementById('database-panel').classList.toggle('locked', !centralAlive);
  document.getElementById('database-lock').style.display = centralAlive ? 'none' : 'block';

  const patientWritable = edgeAlive || centralAlive;
  const disableDbControls = !centralAlive;
  ['btn-med-save', 'btn-stock-add'].forEach(function(id) {
    const el = document.getElementById(id);
    if (el) el.disabled = disableDbControls;
  });
  ['m-id', 'm-name', 'm-unit', 'm-min', 's-med', 's-qty', 's-by'].forEach(function(id) {
    const el = document.getElementById(id);
    if (el) el.disabled = disableDbControls;
  });
  ['btn-patient-save', 'p-id', 'p-name', 'p-dob', 'p-contact', 'p-branch'].forEach(function(id) {
    const el = document.getElementById(id);
    if (el) el.disabled = !patientWritable;
  });

  if (edgeAlive && centralAlive) {
    setBanner('good', 'Both nodes are online. Inventory and dispensing use edge; database view uses central.');
  } else if (edgeAlive && !centralAlive) {
    setBanner('warn', 'Central is down. Edge inventory, dispensing, and patient registration remain available; database view is disabled.');
  } else if (!edgeAlive && centralAlive) {
    setBanner('warn', 'Edge is down. Dispensing is routed through central with branch override.');
  } else {
    setBanner('bad', 'Neither node is reachable right now. Check both VMs and network connectivity.');
  }

  updateNodeCard('edge', edgeState, EDGE_API);
  updateNodeCard('central', centralState, CENTRAL_API);
  updateRoutingSummary();
}

function updateRoutingSummary() {
  const box = document.getElementById('routing-summary');
  box.innerHTML = '';

  const routing = [];
  if (edgeState.alive) {
    routing.push(['Dispense → edge', 'good']);
  } else if (centralState.alive) {
    routing.push(['Dispense → central fallback', 'warn']);
  } else {
    routing.push(['Dispense unavailable', 'bad']);
  }

  routing.push(centralState.alive ? ['Database → central', 'good'] : ['Database disabled', 'bad']);
  routing.push(BRANCH_ID ? ['Branch: ' + BRANCH_ID, 'neutral'] : ['Branch not set', 'warn']);

  routing.forEach(function(item) {
    box.appendChild(makeChip(item[0], item[1]));
  });
}

function openTab(id) {
  if (id === 'database' && !centralState.alive) {
    setBanner('warn', 'Central node is offline, so database actions are disabled.');
  }

  activeTab = id;
  document.querySelectorAll('.panel').forEach(function(el) { el.classList.remove('active'); });
  var panel = document.getElementById('panel-' + id);
  if (panel) panel.classList.add('active');

  document.querySelectorAll('.nav-item').forEach(function(el) {
    el.classList.toggle('active', el.id === 'nav-' + id);
  });

  refreshActiveTab();
}

function currentDispenseBase() {
  if (edgeState.alive) return EDGE_API;
  if (centralState.alive) return CENTRAL_API;
  return '';
}

function applyCentralBranchOverride(payload) {
  payload.branch_override = BRANCH_ID || (centralState.info && centralState.info.branch ? centralState.info.branch : '') || 'central';
  return payload;
}

async function nodeHasStock(base, medicationID, quantity) {
  try {
    const items = await fetchJson(base, '/api/inventory');
    const match = Array.isArray(items) ? items.find(function(item) {
      return item.medication_id === medicationID;
    }) : null;
    const rawBalance = match ? (match.balance != null ? match.balance : match.quantity) : 0;
    return Number(rawBalance || 0) >= quantity;
  } catch (err) {
    return false;
  }
}

async function chooseDispenseBase(payload) {
  if (edgeState.alive && centralState.alive) {
    if (await nodeHasStock(EDGE_API, payload.medication_id, payload.quantity)) {
      return EDGE_API;
    }
    applyCentralBranchOverride(payload);
    return CENTRAL_API;
  }
  if (edgeState.alive) return EDGE_API;
  if (centralState.alive) {
    applyCentralBranchOverride(payload);
    return CENTRAL_API;
  }
  return '';
}

function totalInventoryBalance(items) {
  return Array.isArray(items)
    ? items.reduce(function(sum, item) {
      return sum + Number(item.balance != null ? item.balance : (item.quantity || 0));
    }, 0)
    : 0;
}

async function ensureEdgeSeededFromCentral() {
  if (!edgeState.alive || !centralState.alive || edgeSeedAttempted) return;

  try {
    const results = await Promise.all([
      fetchJson(EDGE_API, '/api/inventory'),
      fetchJson(CENTRAL_API, '/api/inventory'),
    ]);
    const edgeItems = results[0];
    const centralItems = results[1];
    const shouldSeed = Array.isArray(centralItems) && centralItems.length > 0 &&
      (!Array.isArray(edgeItems) || edgeItems.length === 0 ||
        (totalInventoryBalance(centralItems) > 0 && totalInventoryBalance(edgeItems) === 0));
    if (!shouldSeed) return;

    edgeSeedAttempted = true;
    const snapshot = await fetchJson(CENTRAL_API, '/api/sync/export', {}, 6000);
    await fetchJson(EDGE_API, '/api/sync/import', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(snapshot),
    }, 8000);
    setBanner('good', 'Edge operational inventory was seeded from central.');
  } catch (err) {
    edgeSeedAttempted = true;
    setBanner('warn', 'Edge is online, but existing central records could not be seeded to edge: ' + (err.message || err));
  }
}

function currentInventoryBase() {
  if (edgeState.alive) return EDGE_API;
  if (centralState.alive) return CENTRAL_API;
  return '';
}

function currentDatabaseBase() {
  return centralState.alive ? CENTRAL_API : '';
}

async function refreshHealth(forceReloadData) {
  syncSettingsFromInputs();
  updateInputsLive();

  const results = await Promise.all([
    probeStatus(EDGE_API),
    probeStatus(CENTRAL_API),
  ]);
  edgeState = results[0];
  centralState = results[1];

  updateUiState();

  await ensureEdgeSeededFromCentral();

  if (forceReloadData || activeTab === 'dispense' || activeTab === 'inventory' || activeTab === 'database') {
    await refreshActiveTab();
  }
}

async function refreshActiveTab() {
  if (activeTab === 'dispense') {
    await Promise.all([loadInventoryForDispense(), loadRecentTransactions()]);
  } else if (activeTab === 'inventory') {
    await loadInventoryTable();
  } else if (activeTab === 'database') {
    if (centralState.alive) {
      await Promise.all([loadDatabaseSelects(), loadDatabaseView(dbView)]);
    } else {
      document.getElementById('db-view').innerHTML = '<div class="empty">Central node is offline.</div>';
    }
  } else if (activeTab === 'status') {
  }
}

async function loadInventoryForDispense() {
  const base = currentInventoryBase();
  const select = document.getElementById('d-med');
  select.innerHTML = '';

  if (!base) {
    select.innerHTML = '<option value="">No node available</option>';
    return;
  }

  try {
    const items = await fetchJson(base, '/api/inventory');
    document.getElementById('inventory-source').textContent = 'Source: ' + (base === CENTRAL_API ? 'central' : 'edge');

    if (!Array.isArray(items) || items.length === 0) {
      const opt = document.createElement('option');
      opt.value = '';
      opt.textContent = 'No medications available';
      select.appendChild(opt);
      return;
    }

    items.forEach(function(it) {
      const opt = document.createElement('option');
      opt.value = it.medication_id;
      opt.textContent = it.name + ' (' + it.medication_id + ')';
      select.appendChild(opt);
    });
  } catch (err) {
    select.innerHTML = '<option value="">Failed to load</option>';
  }
}

async function loadRecentTransactions() {
  const wrap = document.getElementById('recent-tx');
  const sources = [];
  if (edgeState.alive) sources.push({ label: 'edge', base: EDGE_API });
  if (centralState.alive) sources.push({ label: 'central', base: CENTRAL_API });
  if (!sources.length) {
    wrap.innerHTML = '<div class="empty">No node available.</div>';
    return;
  }

  try {
    const results = await Promise.all(sources.map(function(source) {
      return fetchJson(source.base, '/api/transactions')
        .then(function(items) {
          return { source: source, items: Array.isArray(items) ? items : [] };
        })
        .catch(function(err) {
          return { source: source, err: err };
        });
    }));
    const failed = results.filter(function(result) { return result.err; });
    if (failed.length === results.length) {
      throw failed[0].err;
    }

    const byId = new Map();
    results.forEach(function(result) {
      if (result.err) return;
      result.items.forEach(function(tx) {
        const key = tx.tx_id || (result.source.label + '-' + byId.size);
        if (byId.has(key)) {
          const existing = byId.get(key);
          if (!existing.source.includes(result.source.label)) {
            existing.source += '+' + result.source.label;
          }
          return;
        }
        const row = Object.assign({}, tx);
        row.source = result.source.label;
        byId.set(key, row);
      });
    });
    const txs = Array.from(byId.values()).sort(function(a, b) {
      return String(b.timestamp || '').localeCompare(String(a.timestamp || ''));
    });
    const loadedSources = results
      .filter(function(result) { return !result.err; })
      .map(function(result) { return result.source.label; })
      .join(' + ');
    document.getElementById('recent-source').textContent = failed.length ? ('Source: ' + loadedSources + ' (partial)') : ('Source: ' + loadedSources);
    wrap.innerHTML = '';
    wrap.appendChild(buildTable(txs, [
      { key: 'source', label: 'Source' },
      { key: 'tx_id', label: 'Tx ID' },
      { key: 'medication_id', label: 'Medication' },
      { key: 'type', label: 'Type' },
      { key: 'quantity', label: 'Qty' },
      { key: 'pharmacist_id', label: 'Pharmacist' },
      { key: 'branch_id', label: 'Branch' },
      { key: 'node_id', label: 'Node' },
      { key: 'timestamp', label: 'Time' },
    ]));
  } catch (err) {
    wrap.innerHTML = '<div class="empty">' + esc(err.message) + '</div>';
  }
}

async function loadInventoryTable() {
  const base = currentInventoryBase();
  const wrap = document.getElementById('inventory-table');
  if (!base) {
    wrap.innerHTML = '<div class="empty">No reachable node.</div>';
    return;
  }

  try {
    const items = await fetchJson(base, '/api/inventory');
    document.getElementById('inventory-source').textContent = 'Source: ' + (base === CENTRAL_API ? 'central' : 'edge');
    wrap.innerHTML = '';
    wrap.appendChild(buildTable(items, [
      { key: 'medication_id', label: 'Medication ID' },
      { key: 'name', label: 'Name' },
      { key: 'unit', label: 'Unit' },
      { key: 'min_stock', label: 'Min stock' },
      { key: 'balance', label: 'Balance' },
      { key: 'low_stock', label: 'Low stock' },
    ]));
  } catch (err) {
    wrap.innerHTML = '<div class="empty">' + esc(err.message) + '</div>';
  }
}

async function loadDatabaseSelects() {
  if (!centralState.alive) return;
  try {
    const meds = await fetchJson(CENTRAL_API, '/api/medications');
    const sMed = document.getElementById('s-med');
    const dMed = document.getElementById('d-med');
    sMed.innerHTML = '';
    dMed.innerHTML = '';

    if (!Array.isArray(meds) || meds.length === 0) {
      var emptyOpt = document.createElement('option');
      emptyOpt.value = '';
      emptyOpt.textContent = 'No medications';
      sMed.appendChild(emptyOpt.cloneNode(true));
      dMed.appendChild(emptyOpt);
      return;
    }

    meds.forEach(function(m) {
      const label = m.name + ' (' + m.medication_id + ')';
      const opt1 = document.createElement('option');
      opt1.value = m.medication_id;
      opt1.textContent = label;
      sMed.appendChild(opt1);

      const opt2 = document.createElement('option');
      opt2.value = m.medication_id;
      opt2.textContent = label;
      dMed.appendChild(opt2);
    });
  } catch (err) {
  }
}

async function loadDatabaseView(view) {
  const host = currentDatabaseBase();
  const box = document.getElementById('db-view');
  if (!host) {
    box.innerHTML = '<div class="empty">Central node is offline.</div>';
    return;
  }

  try {
    box.innerHTML = '<div class="empty">Loading…</div>';

    if (view === 'medications') {
      const meds = await fetchJson(host, '/api/medications');
      box.innerHTML = '';
      box.appendChild(buildTable(meds, [
        { key: 'medication_id', label: 'Medication ID' },
        { key: 'name', label: 'Name' },
        { key: 'unit', label: 'Unit' },
        { key: 'min_stock', label: 'Min stock' },
      ]));
    } else if (view === 'patients') {
      const patients = await fetchJson(host, '/api/patients');
      box.innerHTML = '';
      box.appendChild(buildTable(patients, [
        { key: 'patient_id', label: 'Patient ID' },
        { key: 'name', label: 'Name' },
        { key: 'date_of_birth', label: 'DOB' },
        { key: 'contact', label: 'Contact' },
        { key: 'branch_id', label: 'Branch' },
        { key: 'last_updated', label: 'Updated' },
      ]));
    } else if (view === 'transactions') {
      const txs = await fetchJson(host, '/api/transactions');
      box.innerHTML = '';
      box.appendChild(buildTable(txs, [
        { key: 'tx_id', label: 'Tx ID' },
        { key: 'medication_id', label: 'Medication' },
        { key: 'patient_id', label: 'Patient' },
        { key: 'pharmacist_id', label: 'Pharmacist' },
        { key: 'quantity', label: 'Qty' },
        { key: 'branch_id', label: 'Branch' },
        { key: 'node_id', label: 'Node' },
        { key: 'timestamp', label: 'Time' },
        { key: 'status', label: 'Status' },
      ]));
    }
  } catch (err) {
    box.innerHTML = '<div class="empty">' + esc(err.message) + '</div>';
  }
}

function buildTable(items, columns) {
  if (!Array.isArray(items) || items.length === 0) {
    const empty = document.createElement('div');
    empty.className = 'empty';
    empty.textContent = 'No records found.';
    return empty;
  }

  const wrap = document.createElement('div');
  wrap.className = 'table-wrap';

  const table = document.createElement('table');
  const thead = document.createElement('thead');
  const tbody = document.createElement('tbody');
  const trHead = document.createElement('tr');

  columns.forEach(function(col) {
    const th = document.createElement('th');
    th.textContent = col.label;
    trHead.appendChild(th);
  });
  thead.appendChild(trHead);

  items.forEach(function(item) {
    const tr = document.createElement('tr');
    columns.forEach(function(col) {
      const td = document.createElement('td');
      let value = item[col.key];
      if (typeof value === 'boolean') value = value ? 'Yes' : 'No';
      if (value == null || value === '') value = '—';
      td.textContent = String(value);
      tr.appendChild(td);
    });
    tbody.appendChild(tr);
  });

  table.appendChild(thead);
  table.appendChild(tbody);
  wrap.appendChild(table);
  return wrap;
}

async function submitDispense() {
  syncSettingsFromInputs();

  const medId = document.getElementById('d-med').value.trim();
  const qty = parseInt(document.getElementById('d-qty').value, 10);
  const pharm = document.getElementById('d-pharm').value.trim();
  const patient = document.getElementById('d-patient').value.trim();

  if (!medId || !qty || qty <= 0 || !pharm) {
    setBanner('bad', 'Fill in medication, quantity, and pharmacist ID.');
    return;
  }

  const payload = {
    medication_id: medId,
    quantity: qty,
    pharmacist_id: pharm,
    patient_id: patient
  };

  let target = await chooseDispenseBase(payload);
  if (!target) {
    setBanner('bad', 'No node is available for dispensing.');
    return;
  }

  try {
    const res = await fetchJson(target, '/api/dispense', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(payload),
    });

    setBanner('good', 'Dispense recorded on ' + (target === EDGE_API ? 'edge' : 'central fallback') + '. Tx: ' + (res.tx_id || '—'));
    await refreshHealth(false);
  } catch (err) {
    if (target === EDGE_API && centralState.alive && err.status === 409 && String(err.message).includes('insufficient stock')) {
      applyCentralBranchOverride(payload);
      try {
        const res = await fetchJson(CENTRAL_API, '/api/dispense', {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify(payload),
        });
        setBanner('good', 'Dispense recorded on central fallback. Tx: ' + (res.tx_id || '—'));
        await refreshHealth(false);
        return;
      } catch (fallbackErr) {
        setBanner('bad', fallbackErr.message || 'Dispense failed');
        return;
      }
    }
    setBanner('bad', err.message || 'Dispense failed');
  }
}

async function addMedication() {
  if (!centralState.alive) {
    setBanner('warn', 'Central node is offline, so medication registration is disabled.');
    return;
  }

  const med = {
    medication_id: document.getElementById('m-id').value.trim(),
    name: document.getElementById('m-name').value.trim(),
    unit: document.getElementById('m-unit').value.trim() || 'tablet',
    min_stock: parseInt(document.getElementById('m-min').value, 10) || 0
  };

  if (!med.medication_id || !med.name) {
    setBanner('bad', 'Medication ID and name are required.');
    return;
  }

  try {
    await fetchJson(CENTRAL_API, '/api/medications/register', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(med),
    });
    if (edgeState.alive) {
      try {
        await fetchJson(EDGE_API, '/api/medications/register', {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify(med),
        });
        setBanner('good', 'Medication saved on central and edge.');
      } catch (edgeErr) {
        setBanner('warn', 'Medication saved on central, but edge was not updated: ' + edgeErr.message);
      }
    } else {
      setBanner('good', 'Medication saved on central. Edge will need replication before it can dispense it offline.');
    }
    await refreshHealth(true);
    await loadDatabaseView('medications');
  } catch (err) {
    setBanner('bad', err.message || 'Failed to save medication');
  }
}

async function addStock() {
  if (!centralState.alive) {
    setBanner('warn', 'Central node is offline, so stock updates are disabled.');
    return;
  }

  const payload = {
    medication_id: document.getElementById('s-med').value.trim(),
    quantity: parseInt(document.getElementById('s-qty').value, 10),
    added_by: document.getElementById('s-by').value.trim(),
    transaction_id: newClientID('stock-tx'),
    ledger_entry_id: newClientID('stock-ledger'),
    recorded_at: sqlNowUTC(),
    node_id: 'central',
  };

  if (!payload.medication_id || !payload.quantity || payload.quantity <= 0) {
    setBanner('bad', 'Choose a medication and a quantity greater than zero.');
    return;
  }

  try {
    await fetchJson(CENTRAL_API, '/api/stock/add', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(payload),
    });
    if (edgeState.alive) {
      try {
        await fetchJson(EDGE_API, '/api/stock/add', {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify(payload),
        });
        setBanner('good', 'Stock added on central and edge.');
      } catch (edgeErr) {
        setBanner('warn', 'Stock added on central, but edge was not updated: ' + edgeErr.message);
      }
    } else {
      setBanner('good', 'Stock added on central. Edge will need replication before offline dispensing.');
    }
    await refreshHealth(true);
    await loadInventoryTable();
  } catch (err) {
    setBanner('bad', err.message || 'Failed to add stock');
  }
}

async function savePatient() {
  const target = edgeState.alive ? EDGE_API : (centralState.alive ? CENTRAL_API : '');
  if (!target) {
    setBanner('warn', 'No node is available for patient management.');
    return;
  }

  const payload = {
    patient_id: document.getElementById('p-id').value.trim(),
    name: document.getElementById('p-name').value.trim(),
    date_of_birth: document.getElementById('p-dob').value,
    contact: document.getElementById('p-contact').value.trim(),
    branch_id: document.getElementById('p-branch').value.trim() || BRANCH_ID
  };

  if (!payload.patient_id || !payload.name) {
    setBanner('bad', 'Patient ID and name are required.');
    return;
  }

  try {
    await fetchJson(target, '/api/patients/register', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(payload),
    });
    setBanner('good', 'Patient saved on ' + (target === EDGE_API ? 'edge' : 'central') + '.');
    if (centralState.alive) await loadDatabaseView('patients');
  } catch (err) {
    setBanner('bad', err.message || 'Failed to save patient');
  }
}

async function setDatabaseView(view, btn) {
  dbView = view;
  document.querySelectorAll('.tab-btn').forEach(function(el) { el.classList.remove('active'); });
  if (btn) btn.classList.add('active');
  await loadDatabaseView(view);
}

function saveSettings() {
  syncSettingsFromInputs();
  refreshHealth(true);
  setBanner('good', 'Settings saved.');
}

function bindInput(id, key) {
  const el = document.getElementById(id);
  if (!el) return;
  el.addEventListener('input', function() {
    if (key === 'edge') EDGE_API = el.value.trim();
    if (key === 'central') CENTRAL_API = el.value.trim();
    if (key === 'branch') BRANCH_ID = el.value.trim();
    if (key === 'auth') AUTH_TOKEN = normalizeAuthToken(el.value);
    localStorage.setItem(STORAGE.edge, EDGE_API);
    localStorage.setItem(STORAGE.central, CENTRAL_API);
    localStorage.setItem(STORAGE.branch, BRANCH_ID);
    localStorage.setItem(STORAGE.auth, AUTH_TOKEN);
  });
}

function initBindings() {
  bindInput('edge-api', 'edge');
  bindInput('central-api', 'central');
  bindInput('branch-id', 'branch');
  bindInput('auth-token', 'auth');
}

function loadInitialState() {
  loadSettingsFromStorage();
  initBindings();
  updateInputsLive();
}

async function init() {
  loadInitialState();
  await refreshHealth(true);
  if (pollTimer) clearInterval(pollTimer);
  pollTimer = setInterval(function() { refreshHealth(false); }, 3000);
}

init();
</script>
</body>
</html>`

export default function fileImport(file, deviceSerial) {
  return new Promise((resolve, reject) => {
    console.log(file.name, file.type, deviceSerial);
    const reader = new FileReader();

    if (file.type === "text/xml") {
      reader.readAsText(file);
      reader.addEventListener("load", () => {
        const data = reader.result;
        const parser = new DOMParser();
        const xmlDoc = parser.parseFromString(data, "text/xml");
        const result = xmlToJson(xmlDoc);
        Object.assign(dataResult, result);
        resolve(dataResult);
      });
      reader.addEventListener("error", () => {
        reject("File read error");
      });
    } else if (file.name.endsWith(".log")) {
      reader.readAsText(file);
      reader.addEventListener("load", (e) => {
        const data = e.target.result;
        const result = csvToJSON(data);
        Object.assign(dataResult, result);
        resolve(dataResult);
      });
      reader.addEventListener("error", (e) => {
        reject(e);
      });
    } else if (file.name.endsWith(".bnk")) {
      reader.readAsText(file);
      reader.addEventListener("load", (e) => {
        const data = e.target.result;
        const result = bnkToJSON(data);
        Object.assign(dataResult, result);
        resolve(dataResult);
      });
      reader.addEventListener("error", (e) => {
        reject(e);
      });
    } else {
      reject("Unsupported file format");
    }
  });
}

//////////////////////////////////////// XML to JSON Conversion //////////////////////////////////////////
function xmlToJson(xmlDoc) {
  let result = {};

  let dev = evaluateXPath(
    xmlDoc,
    'biotronik-ieee11073-export/dataset/section[@name="MDC"]/section[@name="IDC"]/section[@name="DEV"]//value',
  );
  let sess = evaluateXPath(
    xmlDoc,
    'biotronik-ieee11073-export/dataset/section[@name="MDC"]/section[@name="IDC"]/section[@name="SESS"]//value',
  );
  let stat = evaluateXPath(
    xmlDoc,
    'biotronik-ieee11073-export/dataset/section[@name="MDC"]/section[@name="IDC"]/section[@name="STAT"]//value',
  );

  let msmtBatt = evaluateXPath(
    xmlDoc,
    'biotronik-ieee11073-export/dataset/section[@name="MDC"]/section[@name="IDC"]/section[@name="MSMT"]//section[@name="BATTERY"]//value',
  );
  let msmtCap = evaluateXPath(
    xmlDoc,
    'biotronik-ieee11073-export/dataset/section[@name="MDC"]/section[@name="IDC"]/section[@name="MSMT"]//section[@name="CAP"]//value',
  );
  let msmtRa = evaluateXPath(
    xmlDoc,
    'biotronik-ieee11073-export/dataset/section[@name="MDC"]/section[@name="IDC"]/section[@name="MSMT"]//section[@name="LEADCHNL_RA"]//value',
  );
  let msmtRv = evaluateXPath(
    xmlDoc,
    'biotronik-ieee11073-export/dataset/section[@name="MDC"]/section[@name="IDC"]/section[@name="MSMT"]//section[@name="LEADCHNL_RV"]//value',
  );
  let msmtHv = evaluateXPath(
    xmlDoc,
    'biotronik-ieee11073-export/dataset/section[@name="MDC"]/section[@name="IDC"]/section[@name="MSMT"]//section[@name="LEADHVCHNL"]//value',
  );
  let msmtLv = evaluateXPath(
    xmlDoc,
    'biotronik-ieee11073-export/dataset/section[@name="MDC"]/section[@name="IDC"]/section[@name="MSMT"]//section[@name="LEADCHNL_LV"]//value',
  );
  let brady = evaluateXPath(
    xmlDoc,
    'biotronik-ieee11073-export/dataset/section[@name="MDC"]/section[@name="IDC"]/section[@name="SET"]//section[@name="BRADY"]//value',
  );

  // Parsing XML into respective objects
  // result['dev'] = parseNode(dev);
  // result['sess'] = parseNode(sess);
  // result['msmtBatt'] = parseNode(msmtBatt);
  // result['msmtRa'] = parseNode(msmtRa);
  // result['msmtRv'] = parseNode(msmtRv);
  // result['brady'] = parseNode(brady);
  result["dev"] = dev;
  result["sess"] = sess;
  result["stat"] = stat;
  result["msmtBatt"] = msmtBatt;
  result["msmtCap"] = msmtCap;
  result["msmtRa"] = msmtRa;
  result["msmtRv"] = msmtRv;
  result["msmtHv"] = msmtHv;
  result["msmtLv"] = msmtLv;
  result["brady"] = brady;
  console.log(result);

  // Populate the global dataResult object
  return updateXmlDataResult(result);
}

// Helper function to extract XPath results into an object
function evaluateXPath(xmlDoc, xpath) {
  let result = {};
  let nodes = xmlDoc.evaluate(xpath, xmlDoc, null, XPathResult.ANY_TYPE, null);
  let node;
  while ((node = nodes.iterateNext())) {
    result[node.getAttribute("name")] = node.textContent;
  }
  return result;
}

// Populate dataResult object based on parsed XML
function updateXmlDataResult(parsedData) {
  dataResult.serial = parsedData.dev.SERIAL;
  dataResult.report_date = formatDate(parsedData.sess.DTM);

  // Update dataResult with the XML data
  dataResult.mdc_idc_dev_model = parsedData.dev.MODEL;
  dataResult.mdc_idc_dev_serial = parsedData.dev.SERIAL;
  dataResult.mdc_idc_dev_type = parsedData.dev.TYPE;
  // settings
  dataResult.mdc_idc_set_brady_lowrate = parsedData.brady.LOWRATE;
  dataResult.mdc_idc_dev_sav = parsedData.dev.SAV;
  dataResult.mdc_idc_dev_pav = parsedData.dev.PAV;
  dataResult.mdc_idc_set_brady_max_tracking_rate =
    parsedData.brady.MAX_TRACKING_RATE;
  dataResult.mdc_idc_set_brady_max_sensor_rate =
    parsedData.brady.MAX_SENSOR_RATE;
  dataResult.mdc_idc_set_brady_mode = parsedData.brady.VENDOR_MODE;
  //Battery and capacitor
  dataResult.mdc_idc_batt_status = parsedData.msmtBatt.STATUS;
  dataResult.mdc_idc_batt_volt = parsedData.msmtBatt.VOLTAGE;
  dataResult.mdc_idc_batt_remaining = parsedData.msmtBatt.REMAINING;
  dataResult.mdc_idc_cap_charge_time = parsedData.msmtCap.CHARGE_TIME;
  // Ra Lead Channel
  dataResult.mdc_idc_msmt_ra_impedance_mean = parsedData.msmtRa.VALUE;
  dataResult.mdc_idc_msmt_ra_sensing_mean = parsedData.msmtRa.INTR_AMPL_MEAN;
  dataResult.mdc_idc_msmt_ra_threshold = parsedData.msmtRa.AMPLITUDE;
  dataResult.mdc_idc_msmt_ra_pw = parsedData.msmtRa.PULSEWIDTH;
  // Rv Lead Channel
  dataResult.mdc_idc_msmt_rv_impedance_mean = parsedData.msmtRv.VALUE;
  dataResult.mdc_idc_msmt_rv_sensing_mean = parsedData.msmtRv.INTR_AMPL_MEAN;
  dataResult.mdc_idc_msmt_rv_threshold = parsedData.msmtRv.AMPLITUDE;
  dataResult.mdc_idc_msmt_rv_pw = parsedData.msmtRv.PULSEWIDTH;
  // Hv Lead Channel
  dataResult.mdc_idc_msmt_shock_impedance = parsedData.msmtHv.IMPEDANCE;
  // Lv Lead Channel
  dataResult.mdc_idc_msmt_lv_impedance_mean = parsedData.msmtLv.VALUE;
  // dataResult.mdc_idc_msmt_lv_sensing_mean = parsedData.msmtLv.INTR_AMPL_MEAN;
  dataResult.mdc_idc_msmt_lv_threshold = parsedData.msmtLv.AMPLITUDE;
  dataResult.mdc_idc_msmt_lv_pw = parsedData.msmtLv.PULSEWIDTH;
  // Patient statistics
  dataResult.mdc_idc_stat_brady_ra_percent_paced =
    parsedData.stat.RA_PERCENT_PACED;
  dataResult.mdc_idc_stat_brady_rv_percent_paced =
    parsedData.stat.RV_PERCENT_PACED;
  dataResult.mdc_idc_stat_brady_lv_percent_paced =
    parsedData.stat.LV_PERCENT_PACED;
  dataResult.mdc_idc_stat_ataf_burden_percent = parsedData.stat.BURDEN_PERCENT;
  // dataResult.mdc_idc_stat_shocks_delivered = parsedData.stat.SHOCKS_DELIVERED_RECENT;
  // dataResult.mdc_idc_stae_atp_delivered = parsedData.stat.ATP_DELIVERED_RECENT;
  return dataResult;
}

// Format Date Helper
function formatDate(dtm) {
  return `${dtm.substring(0, 4)}-${dtm.substring(4, 6)}-${dtm.substring(6, 8)}`;
}
////////////////////////////////////////// CSV to JSON Conversion //////////////////////////////////////////
// CSV to JSON Conversion for .log files
function csvToJSON(data) {
  const result = {};
  const FS = String.fromCharCode(28); // File separator; end of file. Or between a concatenation of what might otherwise be separate files.
  const GS = String.fromCharCode(29); // Group separator; between sections of data. Not needed in simple data files.
  const RS = String.fromCharCode(30); // Record separator; end of a record or row.
  const US = String.fromCharCode(31); // Unit separator; between fields of a record, or members of a row.
  let vals = data.split("\n");
  for (let i = 0; i < vals.length; i++) {
    let row = vals[i].split(FS);
    let code = row[0].toString();
    result[code] = { name: row[1], value: row[2] };
  }

  // Assuming the CSV data updates specific parts of dataResult
  // console.log('result', result);
  return updateLogDataResult(result);
}

function updateLogDataResult(parsedData) {
  // report date
  if (parsedData["105"] && parsedData["105"].value) {
    dataResult.report_date = new Date(parsedData[105].value)
      .toISOString()
      .split("T")[0];
  }
  //serial
  if (parsedData["202"] && parsedData["202"].value) {
    dataResult.serial = parsedData["202"].value;
  }
  //model
  if (parsedData["200"] && parsedData["200"].value) {
    dataResult.mdc_idc_dev_model = parsedData["200"].value;
  }
  //type
  if (parsedData["201"] && parsedData["201"].value) {
    dataResult.mdc_idc_dev_type = parsedData["201"].value;
  }
  //mode
  if (parsedData["301"] && parsedData["301"].value) {
    dataResult.mdc_idc_set_brady_mode = parsedData["301"].value;
  }
  //low rate
  if (parsedData["302"] && parsedData["302"].value) {
    dataResult.mdc_idc_set_brady_lowrate = parsedData["302"].value;
  }
  //max tracking rate
  if (parsedData["323"] && parsedData["323"].value) {
    dataResult.mdc_idc_set_brady_max_tracking_rate = parsedData["323"].value;
  }
  //max sensor rate
  if (parsedData["406"] && parsedData["406"].value) {
    dataResult.mdc_idc_set_brady_max_sensor_rate = parsedData["406"].value;
  }
  //battery status
  if (parsedData["519"] && parsedData["519"].value) {
    dataResult.mdc_idc_batt_volt = parsedData["519"].value;
  }
  if (parsedData["533"] && parsedData["533"].value) {
    dataResult.mdc_idc_batt_remaining = parsedData["533"].value;
  }
  if (parsedData["2745"] && parsedData["2745"].value) {
    dataResult.mdc_idc_cap_charge_time = parsedData["2745"].value;
  }
  // patient status
  if (parsedData["2680"] && parsedData["2680"].value) {
    dataResult.mdc_idc_stat_ataf_burden_percent = parsedData["2680"].value;
  }
  if (parsedData["2681"] && parsedData["2681"].value) {
    dataResult.mdc_idc_stat_brady_rv_percent_paced = parsedData["2681"].value;
  }
  if (parsedData["2682"] && parsedData["2682"].value) {
    dataResult.mdc_idc_stat_brady_ra_percent_paced = parsedData["2682"].value;
  }
  // Device measurements
  if (parsedData["2721"] && parsedData["2721"].value) {
    dataResult.mdc_idc_msmt_ra_sensing_mean = parsedData["2721"].value;
  }
  if (parsedData["512"] && parsedData["512"].value) {
    dataResult.mdc_idc_msmt_ra_impedance_mean = parsedData["512"].value;
  }
  if (parsedData["1610"] && parsedData["1610"].value) {
    dataResult.mdc_idc_msmt_ra_threshold = parsedData["1610"].value;
  }
  if (parsedData["1611"] && parsedData["1611"].value) {
    dataResult.mdc_idc_msmt_ra_pw = parsedData["1611"].value;
  }
  if (parsedData["2722"] && parsedData["2722"].value) {
    dataResult.mdc_idc_msmt_rv_sensing_mean = parsedData["2722"].value;
  }
  if (parsedData["507"] && parsedData["507"].value) {
    dataResult.mdc_idc_msmt_rv_impedance_mean = parsedData["507"].value;
  }
  if (parsedData["1620"] && parsedData["1620"].value) {
    dataResult.mdc_idc_msmt_rv_threshold = parsedData["1620"].value;
  }
  if (parsedData["1621"] && parsedData["1621"].value) {
    dataResult.mdc_idc_msmt_rv_pw = parsedData["1621"].value;
  }
  if (parsedData["2730"] && parsedData["2730"].value) {
    dataResult.mdc_idc_msmt_shock_impedance = parsedData["2730"].value;
  }
  if (parsedData["2720"] && parsedData["2720"].value) {
    dataResult.mdc_idc_msmt_lv_impedance_mean = parsedData["2720"].value;
  }
  if (parsedData["1616"] && parsedData["1616"].value) {
    dataResult.mdc_idc_msmt_lv_threshold = parsedData["1616"].value;
  }
  if (parsedData["1617"] && parsedData["1617"].value) {
    dataResult.mdc_idc_msmt_lv_pw = parsedData["1617"].value;
  }
  // Update other relevant fields in dataResult
  return dataResult;
}
////////////////////////////////////////// BNK to JSON Conversion //////////////////////////////////////////
// BNK to JSON Conversion for .bnk files
function bnkToJSON(data) {
  let vals = data.split("\n"); // Split the CSV content into rows
  let result = {}; // Initialize result as an object to store key-value pairs

  for (let i in vals) {
    let row = vals[i].split(",");

    // Skip lines that don't have the expected key-value format
    if (row.length === 2) {
      let key = row[0].trim(); // Trim whitespace from the key
      let value = row[1].trim(); // Trim whitespace from the value

      // Store each key-value pair in the result object
      result[key] = value;
    }
  }

  console.log("result", result); // Output the parsed data for debugging

  // Now update the dataResult object with the parsed data
  return updateBnkDataResult(result);
}

// Function to update the dataResult object with the parsed CSV data
function updateBnkDataResult(parsedData, deviceSerial) {
  // Here are some examples of how to map parsed data to dataResult
  console.log("parsed data", parsedData); // Output the parsed data for debugging
  // Date Conversion
  const implantDay = parsedData["PatientData.ImplantDay"];
  const implantMonth = parsedData["PatientData.ImplantMonth"];
  const numericalMonth = monthToNumber(implantMonth);
  const implantYear = parsedData["PatientData.ImplantYear"];
  dataResult.report_date = `${implantYear}-${numericalMonth}-${implantDay}`;

  // Check if the serial number in the file matches the device serial number
  if (parsedData["SystemSerialNumber"] != deviceSerial) {
    alert(
      "The serial number in the file does not match the device serial number.",
    );
  }
  dataResult.mdc_idc_dev_serial = parsedData["SystemSerialNumber"];
  if (parsedData["SystemName"]) {
    dataResult.mdc_idc_dev_model = parsedData["SystemName"];
  }
  if (parsedData["SystemTypeTierName"]) {
    dataResult.mdc_idc_dev_type = parsedData["SystemTypeTierName"];
  }
  if (parsedData["BdyPSBradyMode"]) {
    dataResult.mdc_idc_set_brady_mode = parsedData["BdyPSBradyMode"];
  }
  if (parsedData["NormParams.LRLIntvl"]) {
    let lrlToBpm = Math.floor(
      60000 / parsedData["NormParams.LRLIntvl"].slice(0, -3),
    );
    dataResult.mdc_idc_set_brady_lowrate = lrlToBpm;
  }
  if (parsedData["NormParams.MTRIntvl"]) {
    dataResult.mdc_idc_set_brady_max_tracking_rate = Math.floor(
      parsedData["NormParams.MTRIntvl"].slice(0, -3),
    );
  }
  if (parsedData["NormParams.MSRIntvl"]) {
    dataResult.mdc_idc_set_brady_max_sensor_rate = Math.floor(
      parsedData["NormParams.MSRIntvl"].slice(0, -3),
    );
  }
  if (parsedData["BatteryStatus.BatteryPhase"]) {
    dataResult.mdc_idc_batt_status = parsedData["BatteryStatus.BatteryPhase"];
  }
  if (parsedData["CapformChargeTime"]) {
    dataResult.mdc_idc_cap_charge_time = parsedData["CapformChargeTime"]
      .slice(0, -2)
      .trim();
  }

  if (parsedData["ManualIntrinsicResult.LVMsmt.Msmt"]) {
    dataResult.mdc_idc_msmt_lv_sensing_mean = parsedData[
      "ManualIntrinsicResult.LVMsmt.Msmt"
    ]
      .slice(0, -3)
      .trim();
  }
  if (parsedData["ManualLeadImpedData.LVMsmt.Msmt"]) {
    dataResult.mdc_idc_msmt_lv_impedance_mean = parsedData[
      "ManualLeadImpedData.LVMsmt.Msmt"
    ]
      .slice(0, -5)
      .trim();
  }
  if (parsedData["InterPaceThreshResult.LVMsmt.Amplitude"]) {
    dataResult.mdc_idc_msmt_lv_threshold = parsedData[
      "InterPaceThreshResult.LVMsmt.Amplitude"
    ]
      .slice(0, -3)
      .trim();
  }
  if (parsedData["InterPaceThreshResult.LVMsmt.PulseWidth"]) {
    dataResult.mdc_idc_msmt_lv_pw = parsedData[
      "InterPaceThreshResult.LVMsmt.PulseWidth"
    ]
      .slice(0, -3)
      .trim();
  }

  if (parsedData["ManualIntrinsicResult.RVMsmt.Msmt"]) {
    dataResult.mdc_idc_msmt_rv_sensing_mean = parsedData[
      "ManualIntrinsicResult.RVMsmt.Msmt"
    ]
      .slice(0, -3)
      .trim();
  }
  if (parsedData["ManualLeadImpedData.RVMsmt.Msmt"]) {
    dataResult.mdc_idc_msmt_rv_impedance_mean = parsedData[
      "ManualLeadImpedData.RVMsmt.Msmt"
    ]
      .slice(0, -5)
      .trim();
  }
  if (parsedData["InterPaceThreshResult.RVMsmt.Amplitude"]) {
    dataResult.mdc_idc_msmt_rv_threshold = parsedData[
      "InterPaceThreshResult.RVMsmt.Amplitude"
    ]
      .slice(0, -3)
      .trim();
  }
  if (parsedData["InterPaceThreshResult.RVMsmt.PulseWidth"]) {
    dataResult.mdc_idc_msmt_rv_pw = parsedData[
      "InterPaceThreshResult.RVMsmt.PulseWidth"
    ]
      .slice(0, -3)
      .trim();
  }
  if (parsedData["ShockLastImpedance"]) {
    dataResult.mdc_idc_msmt_shock_impedance = parsedData["ShockLastImpedance"]
      .slice(0, -5)
      .trim();
  }

  if (parsedData["ManualIntrinsicResult.RAMsmt.Msmt"]) {
    dataResult.mdc_idc_msmt_ra_sensing_mean = parsedData[
      "ManualIntrinsicResult.RAMsmt.Msmt"
    ]
      .slice(0, -3)
      .trim();
  }
  if (parsedData["ManualLeadImpedData.RAMsmt.Msmt"]) {
    dataResult.mdc_idc_msmt_ra_impedance_mean = parsedData[
      "ManualLeadImpedData.RAMsmt.Msmt"
    ]
      .slice(0, -5)
      .trim();
  }
  if (parsedData["InterPaceThreshResult.RAMsmt.Amplitude"]) {
    dataResult.mdc_idc_msmt_ra_threshold = parsedData[
      "InterPaceThreshResult.RAMsmt.Amplitude"
    ]
      .slice(0, -3)
      .trim();
  }
  if (parsedData["InterPaceThreshResult.RAMsmt.PulseWidth"]) {
    dataResult.mdc_idc_msmt_ra_pw = parsedData[
      "InterPaceThreshResult.RAMsmt.PulseWidth"
    ]
      .slice(0, -3)
      .trim();
  }

  // Output the updated dataResult for debugging
  return dataResult;
}

// Initial dataResult template
const dataResult = {
  report_date: "",
  mdc_idc_dev_type: "",
  mdc_idc_dev_model: "",
  mdc_idc_dev_serial: "",
  mdc_idc_stat_ataf_burden_percent: "",
  mdc_idc_stat_pvc_count: "",
  mdc_idc_stat_nsvt_count: "",
  mdc_idc_set_brady_mode: "",
  mdc_idc_set_brady_lowrate: "",
  mdc_idc_set_brady_max_tracking_rate: "",
  mdc_idc_set_brady_max_sensor_rate: "",
  mdc_idc_dev_sav: "",
  mdc_idc_dev_pav: "",
  // mdc_idc_stat_shocks_delivered: "",
  mdc_idc_stat_atp_delivered: "",
  mdc_idc_stat_brady_ra_percent_paced: "",
  mdc_idc_stat_brady_rv_percent_paced: "",
  mdc_idc_stat_brady_lv_percent_paced: "",
  biv_percent_paced: "",
  mdc_idc_batt_volt: "",
  mdc_idc_batt_remaining: "",
  mdc_idc_batt_status: "",
  mdc_idc_cap_charge_time: "",
  mdc_idc_msmt_ra_impedance_mean: "",
  mdc_idc_msmt_ra_sensing_mean: "",
  mdc_idc_msmt_ra_threshold: "",
  mdc_idc_msmt_ra_pw: "",
  mdc_idc_msmt_rv_impedance_mean: "",
  mdc_idc_msmt_rv_sensing_mean: "",
  mdc_idc_msmt_rv_threshold: "",
  mdc_idc_msmt_rv_pw: "",
  mdc_idc_msmt_shock_impedance: "",
  mdc_idc_msmt_lv_impedance_mean: "",
  mdc_idc_msmt_lv_sensing_mean: "",
  mdc_idc_msmt_lv_threshold: "",
  mdc_idc_msmt_lv_pw: "",
  // mdc_idc_stat_heart_rate: "",
};

function monthToNumber(month) {
  const monthMap = {
    Jan: "01",
    Feb: "02",
    Mar: "03",
    Apr: "04",
    May: "05",
    Jun: "06",
    Jul: "07",
    Aug: "08",
    Sep: "09",
    Oct: "10",
    Nov: "11",
    Dec: "12",
  };

  return monthMap[month];
}

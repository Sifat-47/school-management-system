const API = "http://localhost:8080/api";
const $ = id => document.getElementById(id);
// Escape text before putting it into HTML, so user input can't inject markup.
const esc = v => String(v ?? "").replace(/[&<>"']/g,
  c => ({ "&": "&amp;", "<": "&lt;", ">": "&gt;", '"': "&quot;", "'": "&#39;" }[c]));
// Local copy of everything the backend returned.
const data = { students: [], employees: [], classes: [], courses: [] };
let current = "students";
// Look up a teacher's name from their ID (used in the Classes and Courses tables).
const teacherName = id => {
  const t = data.employees.find(e => String(e.id) === String(id));
  return t ? t.name : "";
};
// One config object per section: form fields and table columns.
// The same code below draws all four sections from this.
const SECTIONS = {
  students: {
  label: "Students",
  single: "Student",
  sub: "Everyone enrolled at the school.",
  add: "Add student",
  empty: "No students yet. Add the first one above.",
  defaults: { class_id: "" },
  fields: [
    { key: "first_name", label: "First name" },
    { key: "last_name", label: "Last name" },
    { key: "age", label: "Age", type: "number" },
    { key: "password", label: "Password" },
    { key: "class_id", label: "Class", classSelect: true, optional: true }  // ← New
  ],
  cols: [
    ["ID", s => s.id],
    ["Name", s => `${s.first_name} ${s.last_name}`],
    ["Age", s => s.age],
    ["Class", s => s.class_id || "—"]
  ]
},
  employees: {
    label: "Employees", single: "Employee", sub: "Teachers and staff members.",
    add: "Add employee", empty: "No employees yet. Add a teacher or staff member above.",
    defaults: {},
    fields: [
      { key: "name", label: "Full name" },
      { key: "role", label: "Role", options: ["Teacher", "Staff"] },
      { key: "subject", label: "Subject (teachers only)", optional: true }],
    cols: [["ID", e => e.id], ["Name", e => e.name], ["Role", e => e.role], ["Subject", e => e.subject]]
  },
  classes: {
  label: "Classes",
  single: "Class",
  sub: "Groups of students from Class 6 to Class 12.",
  add: "Add class",
  empty: "No classes yet. Create one above.",
  defaults: { student_ids: [] },
  fields: [
    { key: "name", label: "Class Name (e.g. Class 6 - Section A)" },
    { 
      key: "grade", 
      label: "Grade", 
      options: ["6", "7", "8", "9", "10", "11", "12"] 
    },
    { key: "teacher_id", label: "Teacher", teacher: true, optional: true }
  ],
  cols: [
    ["ID", c => c.id],
    ["Name", c => c.name],
    ["Grade", c => c.grade || "—"],
    ["Teacher", c => teacherName(c.teacher_id)],
    ["Students", c => (c.student_ids || []).length]
  ]
},
  courses: {
    label: "Courses", single: "Course", sub: "Subjects taught at the school.",
    add: "Add course", empty: "No courses yet. Add one above.",
    defaults: {},
    fields: [
      { key: "name", label: "Course name" },
      { key: "description", label: "Description", optional: true },
      { key: "teacher_id", label: "Teacher", teacher: true, optional: true }],
    cols: [["ID", c => c.id], ["Name", c => c.name], ["Description", c => c.description],
           ["Teacher", c => teacherName(c.teacher_id)]]
  }
};
// Talk to the Go backend. Throws if the server answers with an error.
async function api(path, options) {
  const res = await fetch(`${API}/${path}`, options);
  if (!res.ok) throw new Error(`Request failed (${res.status})`);
  return res.status === 204 ? null : res.json().catch(() => null);
}
function toast(message, isError) {
  const el = document.createElement("div");
  el.className = "toast" + (isError ? " err" : "");
  el.textContent = message;
  $("toasts").appendChild(el);
  setTimeout(() => el.remove(), 3200);
}
function setStatus(online) {
  $("status").className = "status " + (online ? "ok" : "bad");
  $("status").textContent = online ? "Backend connected" : "Backend offline";
}
async function loadAll() {
  const names = Object.keys(SECTIONS);
  const results = await Promise.allSettled(names.map(n => api(n)));
  results.forEach((r, i) => { if (r.status === "fulfilled") data[names[i]] = r.value || []; });
  const online = results.some(r => r.status === "fulfilled");
  setStatus(online);
  if (!online) toast("Can't reach the server. Start the Go backend on port 8080.", true);
  render();
}
function fieldHtml(f) {
  const req = f.optional ? "" : "required";

  if (f.options) {
    return `<select name="${f.key}">${f.options.map(o => `<option>${o}</option>`).join("")}</select>`;
  }

  // Teacher dropdown
  if (f.teacher) {
    const teachers = data.employees.filter(e => (e.role || "").toLowerCase() === "teacher");
    return `<select name="${f.key}">
      <option value="">No teacher yet</option>
      ${teachers.map(t => `<option value="${esc(t.id)}">${esc(t.name)}</option>`).join("")}
    </select>`;
  }

  // Class dropdown (NEW)
  if (f.classSelect) {
    return `<select name="${f.key}">
      <option value="">Select Class</option>
      ${data.classes.map(c => `<option value="${esc(c.id)}">${esc(c.name)} (Grade ${c.grade || "?"})</option>`).join("")}
    </select>`;
  }

  return `<input name="${f.key}" type="${f.type || "text"}" ${req} ${f.type === "number" ? 'min="1"' : ""}>`;
}
function render() {
  const s = SECTIONS[current];
  $("title").textContent = s.label;
  $("sub").textContent = s.sub;
  $("nav").innerHTML = Object.entries(SECTIONS).map(([k, v]) =>
    `<button data-s="${k}" aria-current="${k === current}">${v.label}<b>${data[k].length}</b></button>`).join("");
  $("form").innerHTML = s.fields.map(f => `<label>${f.label}${fieldHtml(f)}</label>`).join("") +
    `<button class="btn">${s.add}</button>`;
  renderTable();
}
function renderTable() {
  const s = SECTIONS[current];
  const q = $("search").value.trim().toLowerCase();
  const rows = data[current].filter(r => s.cols.map(c => c[1](r)).join(" ").toLowerCase().includes(q));
  $("thead").innerHTML = `<tr>${s.cols.map(c => `<th>${c[0]}</th>`).join("")}<th></th></tr>`;
  $("tbody").innerHTML = rows.map(r =>
    `<tr>${s.cols.map((c, i) => `<td${i === 0 ? ' class="id"' : ""}>${esc(c[1](r)) || "—"}</td>`).join("")}` +
    `<td><button class="del" data-id="${esc(r.id)}">Delete</button></td></tr>`).join("");
  $("empty").hidden = rows.length > 0;
  $("empty").textContent = q ? "No matches for that search." : s.empty;
}
// Switch section from the sidebar
$("nav").onclick = e => {
  const b = e.target.closest("button");
  if (!b) return;
  current = b.dataset.s;
  $("search").value = "";
  render();
};
$("search").oninput = renderTable;
// Add a new record
$("form").onsubmit = async e => {
  e.preventDefault();
  const name = current, s = SECTIONS[name];
  const form = new FormData(e.target);
  const body = { ...s.defaults };
  for (const f of s.fields) {
    const v = (form.get(f.key) || "").trim();
    if (!v && !f.optional) return toast(`${f.label} is required`, true);
    if (f.key === "grade") {
  body[f.key] = Number(v);
} else {
  body[f.key] = f.type === "number" ? Number(v) : v;
}
  }
  try {
    await api(name, { method: "POST", headers: { "Content-Type": "application/json" }, body: JSON.stringify(body) });
    data[name] = (await api(name)) || [];
    render();
    toast(`${s.single} added`);
  } catch {
    toast(`Could not add ${s.single.toLowerCase()}. Is the backend running?`, true);
  }
};
// Delete: first click asks for confirmation, second click deletes
$("tbody").onclick = async e => {
  const b = e.target.closest(".del");
  if (!b) return;
  if (!b.classList.contains("sure")) {
    b.classList.add("sure");
    b.textContent = "Confirm delete";
    setTimeout(() => { b.classList.remove("sure"); b.textContent = "Delete"; }, 3000);
    return;
  }
  const name = current, s = SECTIONS[name];
  try {
    await api(`${name}/${b.dataset.id}`, { method: "DELETE" });
    data[name] = (await api(name)) || [];
    render();
    toast(`${s.single} deleted`);
  } catch {
    toast(`Could not delete ${s.single.toLowerCase()}.`, true);
  }
};
loadAll();
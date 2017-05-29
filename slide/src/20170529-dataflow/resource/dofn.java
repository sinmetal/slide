// Entityを受け取り、TableRowを返すExample
public static class EntityToTableRowFn extends DoFn<Entity, TableRow> {
	@Override
	public void processElement(ProcessContext c) {

		TableRow row = new TableRow();
        // c.element()でInput Valueを受け取る
		row.set("__key__", c.element().getKey().toString());
		...
        // c.output()にOutput Valueを渡す
		c.output(row);
	}
}
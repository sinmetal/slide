// Entityの配列を受け取って、TableRowの配列を返すExample
public static class EntityToTableRow 
    extends PTransform<PCollection<Entity>,PCollection<TableRow>> {

	@Override
	public PCollection<TableRow> apply(PCollection<Entity> entities) {

		PCollection<TableRow> rows = entities.apply(ParDo.of(new EntityToTableRowFn()));

		return rows;
	}
}
<template>
  <section>
    <!-- filters -->
    <div>
      <span>Income</span>
      <vee-form>
        <vee-field name="category" as="select">
          <option>Bussiness</option>
          <option>Enterteimant</option>
          <option>Others</option>
        </vee-field>
        <vee-field name="filter" as="select">
          <option>Date</option>
          <option>Value</option>
        </vee-field>
        <app-button text="Filter" type="submit" />
      </vee-form>
    </div>
    <!-- table -->
    <div></div>
  </section>
</template>

<script>
export default {
  name: 'Income',
};
</script>

import { PanelBuilders, SceneQueryRunner } from '@grafana/scenes';
import { HeatmapColorMode } from 'app/plugins/panel/heatmap/types';

import { KEY_SQR_METRIC_VIZ_QUERY, trailDS } from '../../shared';
import { AutoQueryDef } from '../types';

export function heatmapGraphBuilder(def: AutoQueryDef) {
  return PanelBuilders.heatmap()
    .setTitle(def.title)
    .setUnit(def.unit)
    .setOption('calculate', false)
    .setOption('color', {
      mode: HeatmapColorMode.Scheme,
      exponent: 0.5,
      scheme: 'Spectral',
      steps: 32,
      reverse: false,
    })
    .setData(
      new SceneQueryRunner({
        key: KEY_SQR_METRIC_VIZ_QUERY,
        datasource: trailDS,
        queries: def.queries,
      })
    );
}

import os
import sys
import json
import argparse
from typing import Dict, Any

from qgis.core import (
    QgsApplication, QgsProject, QgsPrintLayout,
    QgsLayoutItemMap, QgsLayoutSize, QgsUnitTypes,
    QgsRectangle, QgsLayoutPoint, QgsLayoutExporter,
    QgsLayoutRenderContext)
from PyQt5.QtCore import QSize


def fix_geometry(projectPath: str, algorithm: str, parameters: dict) -> str:
    sys.path.append('/usr/share/qgis/python/plugins')
    try:
        QgsApplication.setPrefixPath("/usr", True)
        qgs = QgsApplication([], False)
        qgs.initQgis()
        from qgis import processing
        from processing.core.Processing import Processing
        Processing.initialize()
    except Exception as e:
        qgs.exitQgis()
        qgs.exit()
        print(f'QGIS init error at {projectPath}: {e}', file=sys.stderr)
        return ''

    try:
        if len(projectPath) != 0:
            project = QgsProject.instance()
            project.read(projectPath)
            print(f'QGIS read project {project.fileName()}')
    except Exception as e:
        qgs.exitQgis()
        qgs.exit()
        print(f'QGIS read project error at {projectPath}: {e}', file=sys.stderr)
        return ''

    try:
        o = processing.run(algorithm, parameters)
    except Exception as e:
        print(f'QGIS run error at {projectPath}: {e}', file=sys.stderr)
    qgs.exitQgis()
    return o


def export_image_tile(projectPath: str, block_per_tile: int,
                     degree_per_tile: int, xMin: int, xMax: int,
                     yMin: int, yMax: int, outputPath: str,
                     layerOutputName: str, Layers: tuple[str]) -> None:
    sys.path.append('/usr/share/qgis/python/plugins')
    try:
        QgsApplication.setPrefixPath("/usr", True)
        qgs = QgsApplication([], False)
        qgs.initQgis()
    except Exception as e:
        qgs.exitQgis()
        print(f'QGIS init error at {projectPath}: {e}', file=sys.stderr)
        return

    try:
        project = QgsProject.instance()
        project.read(projectPath)
        print(f'QGIS read project {project.fileName()}')
    except Exception as e:
        qgs.exitQgis()
        print(f'QGIS read project error at {projectPath}: {e}', file=sys.stderr)
        return

    def uncheckAllLayers(project):
        alllayers = []
        for alllayer in project.mapLayers().values():
            alllayers.append(alllayer)
        root = project.layerTreeRoot()
        for alllayer in alllayers:
            node = root.findLayer(alllayer.id())
            node.setItemVisibilityChecked(False)

    def selectNodes(project, Layers: tuple[str]):
        layers = []
        for layer in project.mapLayers().values():
            for LayerName in Layers:
                if layer.name().startswith(LayerName):
                    layers.append(layer)
        root = project.layerTreeRoot()
        for layer in layers:
            node = root.findLayer(layer.id())
            node.setItemVisibilityChecked(True)

    try:
        uncheckAllLayers(project)
        selectNodes(project, Layers)

        layout = QgsPrintLayout(project)
        layout.initializeDefaults()

        pages = layout.pageCollection()
        pages.page(0).setPageSize(
            QgsLayoutSize(block_per_tile, block_per_tile,
                         QgsUnitTypes.LayoutPixels))

        map = QgsLayoutItemMap(layout)
        map.setRect(0, 0, block_per_tile, block_per_tile)
        map.setExtent(QgsRectangle(xMin, yMin, xMax, yMax))
        map.attemptMove(QgsLayoutPoint(0, 0, QgsUnitTypes.LayoutPixels))
        map.attemptResize(QgsLayoutSize(
            block_per_tile, block_per_tile, QgsUnitTypes.LayoutPixels))
        layout.addLayoutItem(map)

        exporter = QgsLayoutExporter(layout)
        settings = QgsLayoutExporter.ImageExportSettings()
        settings.imageSize = (QSize(block_per_tile, block_per_tile))

        context = QgsLayoutRenderContext(layout)
        context.setFlag(context.FlagAntialiasing, False)
        settings.flags = context.flags()

        ret = exporter.exportToImage(outputPath, settings)
        if ret != 0:
            print(f"exportToImage error: {ret}", file=sys.stderr)
            raise Exception(f"Export failed with return code {ret}")
            
        print(f"Generated {os.path.basename(outputPath)}")

        # Clean up
        del settings
        del exporter
        del context
        layout.removeLayoutItem(map)
        del map
        del pages
        del layout
        uncheckAllLayers(project)

    except Exception as e:
        print(f'QGIS export image error: {e}', file=sys.stderr)
        raise
    finally:
        del project
        qgs.exitQgis()


def main():
    parser = argparse.ArgumentParser(description='QGIS Controller')
    subparsers = parser.add_subparsers(dest='command', help='Commands')

    # Fix geometry command
    fix_parser = subparsers.add_parser('fix_geometry', help='Fix geometry')
    fix_parser.add_argument('--project-path', required=True)
    fix_parser.add_argument('--algorithm', required=True)
    fix_parser.add_argument('--parameters', required=True, type=str,
                          help='Parameters as a JSON string')

    # Export image tile command
    export_parser = subparsers.add_parser('export_image_tile', help='Export image tile')
    export_parser.add_argument('--project-path', required=True)
    export_parser.add_argument('--block-per-tile', type=int, required=True)
    export_parser.add_argument('--degree-per-tile', type=int, required=True)
    export_parser.add_argument('--x-min', type=int, required=True)
    export_parser.add_argument('--x-max', type=int, required=True)
    export_parser.add_argument('--y-min', type=int, required=True)
    export_parser.add_argument('--y-max', type=int, required=True)
    export_parser.add_argument('--output-path', required=True)
    export_parser.add_argument('--layer-output-name', required=True)
    export_parser.add_argument('--layers', required=True,
                             help='Layers as a Python tuple string')

    args = parser.parse_args()

    if args.command == 'fix_geometry':
        parameters = eval(args.parameters)  # Convert string to dict
        fix_geometry(args.project_path, args.algorithm, parameters)
    elif args.command == 'export_image_tile':
        layers = eval(args.layers)  # Convert string to tuple
        export_image_tile(args.project_path, args.block_per_tile,
                         args.degree_per_tile, args.x_min,
                         args.x_max, args.y_min,
                         args.y_max, args.output_path,
                         args.layer_output_name, layers)


if __name__ == '__main__':
    main()
